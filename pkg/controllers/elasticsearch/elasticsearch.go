package elasticsearch

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	appsinformers "k8s.io/client-go/informers/apps/v1beta1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1beta1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	"github.com/jetstack/navigator/pkg/apis/navigator"
	"github.com/jetstack/navigator/pkg/apis/navigator/v1alpha1"
	clientset "github.com/jetstack/navigator/pkg/client/clientset/versioned"
	informerv1alpha1 "github.com/jetstack/navigator/pkg/client/informers/externalversions/navigator/v1alpha1"
	listersv1alpha1 "github.com/jetstack/navigator/pkg/client/listers/navigator/v1alpha1"
	"github.com/jetstack/navigator/pkg/controllers"
	"github.com/jetstack/navigator/pkg/controllers/elasticsearch/configmap"
	"github.com/jetstack/navigator/pkg/controllers/elasticsearch/nodepool"
	"github.com/jetstack/navigator/pkg/controllers/elasticsearch/service"
	"github.com/jetstack/navigator/pkg/controllers/elasticsearch/serviceaccount"
)

type ElasticsearchController struct {
	kubeClient      kubernetes.Interface
	navigatorClient clientset.Interface

	esLister       listersv1alpha1.ElasticsearchClusterLister
	esListerSynced cache.InformerSynced

	pilotLister       listersv1alpha1.PilotLister
	pilotListerSynced cache.InformerSynced

	statefulSetLister       appslisters.StatefulSetLister
	statefulSetListerSynced cache.InformerSynced

	podLister       corelisters.PodLister
	podListerSynced cache.InformerSynced

	serviceAccountLister       corelisters.ServiceAccountLister
	serviceAccountListerSynced cache.InformerSynced

	serviceLister       corelisters.ServiceLister
	serviceListerSynced cache.InformerSynced

	configMapLister       corelisters.ConfigMapLister
	configMapListerSynced cache.InformerSynced

	queue                       workqueue.RateLimitingInterface
	elasticsearchClusterControl ControlInterface
	recorder                    record.EventRecorder
}

// NewElasticsearch returns a new ElasticsearchController that can be used
// to monitor for Elasticsearch resources and create clusters in a target Kubernetes
// cluster.
//
// It accepts a list of informers that are then used to monitor the state of the
// target cluster.
func NewElasticsearch(
	es cache.SharedIndexInformer,
	pilots cache.SharedIndexInformer,
	statefulsets cache.SharedIndexInformer,
	pods cache.SharedIndexInformer,
	serviceaccounts cache.SharedIndexInformer,
	services cache.SharedIndexInformer,
	configmaps cache.SharedIndexInformer,
	cl kubernetes.Interface,
	navigatorCl clientset.Interface,
	recorder record.EventRecorder,
) *ElasticsearchController {
	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "elasticsearchCluster")
	// create a new ElasticsearchController to manage ElasticsearchCluster resources
	elasticsearchController := &ElasticsearchController{
		kubeClient:      cl,
		navigatorClient: navigatorCl,
		queue:           queue,
		recorder:        recorder,
	}

	// add an event handler to the ElasticsearchCluster informer
	es.AddEventHandler(&controllers.QueuingEventHandler{Queue: queue})
	elasticsearchController.esLister = listersv1alpha1.NewElasticsearchClusterLister(es.GetIndexer())
	elasticsearchController.esListerSynced = es.HasSynced

	// add an event handler to the Pilot informer
	pilots.AddEventHandler(&controllers.BlockingEventHandler{WorkFunc: elasticsearchController.handleObject})
	elasticsearchController.pilotLister = listersv1alpha1.NewPilotLister(pilots.GetIndexer())
	elasticsearchController.pilotListerSynced = pilots.HasSynced

	// add an event handler to the StatefulSet informer
	statefulsets.AddEventHandler(&controllers.BlockingEventHandler{WorkFunc: elasticsearchController.handleObject})
	elasticsearchController.statefulSetLister = appslisters.NewStatefulSetLister(statefulsets.GetIndexer())
	elasticsearchController.statefulSetListerSynced = statefulsets.HasSynced

	// add an event handler to the Pod informer
	pods.AddEventHandler(&controllers.BlockingEventHandler{WorkFunc: elasticsearchController.handlePodObject})
	elasticsearchController.podLister = corelisters.NewPodLister(pods.GetIndexer())
	elasticsearchController.podListerSynced = pods.HasSynced

	// add an event handler to the ServiceAccount informer
	serviceaccounts.AddEventHandler(&controllers.BlockingEventHandler{WorkFunc: elasticsearchController.handleObject})
	elasticsearchController.serviceAccountLister = corelisters.NewServiceAccountLister(serviceaccounts.GetIndexer())
	elasticsearchController.serviceAccountListerSynced = serviceaccounts.HasSynced

	// add an event handler to the Service informer
	services.AddEventHandler(&controllers.BlockingEventHandler{WorkFunc: elasticsearchController.handleObject})
	elasticsearchController.serviceLister = corelisters.NewServiceLister(services.GetIndexer())
	elasticsearchController.serviceListerSynced = services.HasSynced

	// add an event handler to the Service informer
	configmaps.AddEventHandler(&controllers.BlockingEventHandler{WorkFunc: elasticsearchController.handleObject})
	elasticsearchController.configMapLister = corelisters.NewConfigMapLister(configmaps.GetIndexer())
	elasticsearchController.configMapListerSynced = configmaps.HasSynced

	// create the actual ElasticsearchCluster controller
	elasticsearchController.elasticsearchClusterControl = NewController(
		elasticsearchController.kubeClient,
		elasticsearchController.navigatorClient,
		elasticsearchController.statefulSetLister,
		elasticsearchController.serviceAccountLister,
		elasticsearchController.serviceLister,
		nodepool.NewController(
			cl,
			navigatorCl,
			elasticsearchController.statefulSetLister,
			elasticsearchController.podLister,
			elasticsearchController.pilotLister,
			recorder,
		),
		configmap.NewController(
			cl,
			elasticsearchController.configMapLister,
			recorder,
		),
		serviceaccount.NewController(
			cl,
			elasticsearchController.serviceAccountLister,
			recorder,
		),
		service.NewController(
			cl,
			elasticsearchController.serviceLister,
			recorder,
		),
		recorder,
	)

	return elasticsearchController
}

// Run is the main event loop
func (e *ElasticsearchController) Run(workers int, stopCh <-chan struct{}) error {
	glog.Infof("Starting Elasticsearch controller")

	if !cache.WaitForCacheSync(stopCh, e.esListerSynced, e.pilotListerSynced, e.statefulSetListerSynced, e.podListerSynced, e.serviceAccountListerSynced, e.serviceListerSynced) {
		return fmt.Errorf("timed out waiting for caches to sync")
	}

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			wait.Until(e.worker, time.Second, stopCh)
		}()
	}

	<-stopCh
	e.queue.ShutDown()
	glog.V(4).Infof("Shutting down Elasticsearch controller workers...")
	wg.Wait()
	glog.V(4).Infof("Elasticsearch controller workers stopped.")
	return nil
}

func (e *ElasticsearchController) worker() {
	glog.V(4).Infof("start worker loop")
	for e.processNextWorkItem() {
		glog.V(4).Infof("processed work item")
	}
	glog.V(4).Infof("exiting worker loop")
}

func (e *ElasticsearchController) processNextWorkItem() bool {
	key, quit := e.queue.Get()
	if quit {
		return false
	}
	defer e.queue.Done(key)

	if k, ok := key.(string); ok {
		if err := e.sync(k); err != nil {
			glog.Infof("Error syncing ElasticsearchCluster %v, requeuing: %v", key.(string), err)
			e.queue.AddRateLimited(key)
		} else {
			e.queue.Forget(key)
		}
	}

	return true
}

func (e *ElasticsearchController) sync(key string) (err error) {
	startTime := time.Now()
	defer func() {
		glog.Infof("Finished syncing elasticsearchcluster %q (%v)", key, time.Now().Sub(startTime))
	}()

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	es, err := e.esLister.ElasticsearchClusters(namespace).Get(name)
	if errors.IsNotFound(err) {
		glog.Infof("ElasticsearchCluster has been deleted %v", key)
		return nil
	}
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to retrieve ElasticsearchCluster %v from store: %v", key, err))
		return err
	}

	status, err := e.elasticsearchClusterControl.Sync(es)

	updateErr := e.updateStatus(es, status)
	if err == nil {
		return updateErr
	}

	return utilerrors.NewAggregate([]error{err, updateErr})
}

func (e *ElasticsearchController) updateStatus(c *v1alpha1.ElasticsearchCluster, status v1alpha1.ElasticsearchClusterStatus) error {
	copy := c.DeepCopy()
	copy.Status = status
	// todo: replace with UpdateStatus
	_, err := e.navigatorClient.NavigatorV1alpha1().ElasticsearchClusters(c.Namespace).UpdateStatus(copy)
	return err
}

func (e *ElasticsearchController) enqueueElasticsearchCluster(obj interface{}) {
	key, err := controllers.KeyFunc(obj)
	if err != nil {
		glog.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	glog.V(4).Infof("Adding ES Cluster '%s' to queue", key)
	e.queue.AddRateLimited(key)
}

func (e *ElasticsearchController) handleObject(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		glog.Errorf("error decoding object, invalid type")
		return
	}
	glog.V(4).Infof("Processing object: %s", object.GetName())
	ownerRef := metav1.GetControllerOf(object)
	if ownerRef == nil || ownerRef.Kind != "ElasticsearchCluster" {
		return
	}

	cluster, err := e.esLister.ElasticsearchClusters(object.GetNamespace()).Get(ownerRef.Name)
	if err != nil {
		glog.V(4).Infof("ignoring orphaned object '%s' of elasticsearchcluster '%s'", object.GetSelfLink(), ownerRef.Name)
		return
	}

	e.enqueueElasticsearchCluster(cluster)
}

// getPodOwner will return the owning ElasticsearchCluster for a pod by
// first looking up it's owning StatefulSet, and then finding the
// ElasticsearchCluster that owns that StatefulSet. If the pod is not managed
// by a StatefulSet/ElasticsearchCluster, it will do nothing.
func (e *ElasticsearchController) handlePodObject(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		glog.Errorf("error decoding object, invalid type")
		return
	}
	glog.V(4).Infof("Processing object: %s", object.GetName())
	ownerRef := metav1.GetControllerOf(object)
	if ownerRef == nil || ownerRef.Kind != "StatefulSet" {
		return
	}

	ss, err := e.statefulSetLister.StatefulSets(object.GetNamespace()).Get(ownerRef.Name)
	if err != nil {
		glog.V(4).Infof("ignoring orphaned object '%s' of statefulset '%s'", object.GetSelfLink(), ownerRef.Name)
		return
	}

	e.handleObject(ss)
}

func init() {
	controllers.Register("ElasticSearch", func(ctx *controllers.Context) controllers.Interface {
		e := NewElasticsearch(
			ctx.SharedInformerFactory.InformerFor(
				ctx.Namespace,
				metav1.GroupVersionKind{Group: navigator.GroupName, Version: "v1alpha1", Kind: "ElasticsearchCluster"},
				informerv1alpha1.NewElasticsearchClusterInformer(ctx.NavigatorClient, ctx.Namespace, time.Second*30, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}),
			),
			ctx.SharedInformerFactory.InformerFor(
				ctx.Namespace,
				metav1.GroupVersionKind{Group: navigator.GroupName, Version: "v1alpha1", Kind: "Pilot"},
				informerv1alpha1.NewPilotInformer(ctx.NavigatorClient, ctx.Namespace, time.Second*30, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}),
			),
			ctx.SharedInformerFactory.InformerFor(
				ctx.Namespace,
				metav1.GroupVersionKind{Group: "apps", Version: "v1beta1", Kind: "StatefulSet"},
				appsinformers.NewStatefulSetInformer(ctx.Client, ctx.Namespace, time.Second*30, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}),
			),
			ctx.SharedInformerFactory.InformerFor(
				ctx.Namespace,
				metav1.GroupVersionKind{Version: "v1", Kind: "Pod"},
				coreinformers.NewPodInformer(ctx.Client, ctx.Namespace, time.Second*30, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}),
			),
			ctx.SharedInformerFactory.InformerFor(
				ctx.Namespace,
				metav1.GroupVersionKind{Version: "v1", Kind: "ServiceAccount"},
				coreinformers.NewServiceAccountInformer(ctx.Client, ctx.Namespace, time.Second*30, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}),
			),
			ctx.SharedInformerFactory.InformerFor(
				ctx.Namespace,
				metav1.GroupVersionKind{Version: "v1", Kind: "Service"},
				coreinformers.NewServiceInformer(ctx.Client, ctx.Namespace, time.Second*30, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}),
			),
			ctx.SharedInformerFactory.InformerFor(
				ctx.Namespace,
				metav1.GroupVersionKind{Version: "v1", Kind: "ConfigMap"},
				coreinformers.NewConfigMapInformer(ctx.Client, ctx.Namespace, time.Second*30, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}),
			),
			ctx.Client,
			ctx.NavigatorClient,
			ctx.Recorder,
		)

		return e.Run
	})
}
