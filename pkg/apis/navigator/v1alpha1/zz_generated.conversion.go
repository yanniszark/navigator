// +build !ignore_autogenerated

/*
Copyright 2017 Jetstack Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	navigator "github.com/jetstack-experimental/navigator/pkg/apis/navigator"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_CouchbaseCluster_To_navigator_CouchbaseCluster,
		Convert_navigator_CouchbaseCluster_To_v1alpha1_CouchbaseCluster,
		Convert_v1alpha1_CouchbaseClusterList_To_navigator_CouchbaseClusterList,
		Convert_navigator_CouchbaseClusterList_To_v1alpha1_CouchbaseClusterList,
		Convert_v1alpha1_CouchbaseClusterNodePool_To_navigator_CouchbaseClusterNodePool,
		Convert_navigator_CouchbaseClusterNodePool_To_v1alpha1_CouchbaseClusterNodePool,
		Convert_v1alpha1_CouchbaseClusterPersistenceConfig_To_navigator_CouchbaseClusterPersistenceConfig,
		Convert_navigator_CouchbaseClusterPersistenceConfig_To_v1alpha1_CouchbaseClusterPersistenceConfig,
		Convert_v1alpha1_CouchbaseClusterPlugin_To_navigator_CouchbaseClusterPlugin,
		Convert_navigator_CouchbaseClusterPlugin_To_v1alpha1_CouchbaseClusterPlugin,
		Convert_v1alpha1_CouchbaseClusterSpec_To_navigator_CouchbaseClusterSpec,
		Convert_navigator_CouchbaseClusterSpec_To_v1alpha1_CouchbaseClusterSpec,
		Convert_v1alpha1_CouchbaseClusterStatus_To_navigator_CouchbaseClusterStatus,
		Convert_navigator_CouchbaseClusterStatus_To_v1alpha1_CouchbaseClusterStatus,
		Convert_v1alpha1_CouchbaseImage_To_navigator_CouchbaseImage,
		Convert_navigator_CouchbaseImage_To_v1alpha1_CouchbaseImage,
		Convert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster,
		Convert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster,
		Convert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList,
		Convert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList,
		Convert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool,
		Convert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool,
		Convert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig,
		Convert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig,
		Convert_v1alpha1_ElasticsearchClusterPlugin_To_navigator_ElasticsearchClusterPlugin,
		Convert_navigator_ElasticsearchClusterPlugin_To_v1alpha1_ElasticsearchClusterPlugin,
		Convert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec,
		Convert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec,
		Convert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus,
		Convert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus,
		Convert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage,
		Convert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage,
	)
}

func autoConvert_v1alpha1_CouchbaseCluster_To_navigator_CouchbaseCluster(in *CouchbaseCluster, out *navigator.CouchbaseCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_CouchbaseClusterSpec_To_navigator_CouchbaseClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CouchbaseClusterStatus_To_navigator_CouchbaseClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_CouchbaseCluster_To_navigator_CouchbaseCluster is an autogenerated conversion function.
func Convert_v1alpha1_CouchbaseCluster_To_navigator_CouchbaseCluster(in *CouchbaseCluster, out *navigator.CouchbaseCluster, s conversion.Scope) error {
	return autoConvert_v1alpha1_CouchbaseCluster_To_navigator_CouchbaseCluster(in, out, s)
}

func autoConvert_navigator_CouchbaseCluster_To_v1alpha1_CouchbaseCluster(in *navigator.CouchbaseCluster, out *CouchbaseCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_navigator_CouchbaseClusterSpec_To_v1alpha1_CouchbaseClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_navigator_CouchbaseClusterStatus_To_v1alpha1_CouchbaseClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_navigator_CouchbaseCluster_To_v1alpha1_CouchbaseCluster is an autogenerated conversion function.
func Convert_navigator_CouchbaseCluster_To_v1alpha1_CouchbaseCluster(in *navigator.CouchbaseCluster, out *CouchbaseCluster, s conversion.Scope) error {
	return autoConvert_navigator_CouchbaseCluster_To_v1alpha1_CouchbaseCluster(in, out, s)
}

func autoConvert_v1alpha1_CouchbaseClusterList_To_navigator_CouchbaseClusterList(in *CouchbaseClusterList, out *navigator.CouchbaseClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]navigator.CouchbaseCluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_CouchbaseClusterList_To_navigator_CouchbaseClusterList is an autogenerated conversion function.
func Convert_v1alpha1_CouchbaseClusterList_To_navigator_CouchbaseClusterList(in *CouchbaseClusterList, out *navigator.CouchbaseClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CouchbaseClusterList_To_navigator_CouchbaseClusterList(in, out, s)
}

func autoConvert_navigator_CouchbaseClusterList_To_v1alpha1_CouchbaseClusterList(in *navigator.CouchbaseClusterList, out *CouchbaseClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]CouchbaseCluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_navigator_CouchbaseClusterList_To_v1alpha1_CouchbaseClusterList is an autogenerated conversion function.
func Convert_navigator_CouchbaseClusterList_To_v1alpha1_CouchbaseClusterList(in *navigator.CouchbaseClusterList, out *CouchbaseClusterList, s conversion.Scope) error {
	return autoConvert_navigator_CouchbaseClusterList_To_v1alpha1_CouchbaseClusterList(in, out, s)
}

func autoConvert_v1alpha1_CouchbaseClusterNodePool_To_navigator_CouchbaseClusterNodePool(in *CouchbaseClusterNodePool, out *navigator.CouchbaseClusterNodePool, s conversion.Scope) error {
	out.Name = in.Name
	out.Replicas = in.Replicas
	out.Roles = *(*[]string)(unsafe.Pointer(&in.Roles))
	out.Resources = (*v1.ResourceRequirements)(unsafe.Pointer(in.Resources))
	out.Persistence = (*navigator.CouchbaseClusterPersistenceConfig)(unsafe.Pointer(in.Persistence))
	return nil
}

// Convert_v1alpha1_CouchbaseClusterNodePool_To_navigator_CouchbaseClusterNodePool is an autogenerated conversion function.
func Convert_v1alpha1_CouchbaseClusterNodePool_To_navigator_CouchbaseClusterNodePool(in *CouchbaseClusterNodePool, out *navigator.CouchbaseClusterNodePool, s conversion.Scope) error {
	return autoConvert_v1alpha1_CouchbaseClusterNodePool_To_navigator_CouchbaseClusterNodePool(in, out, s)
}

func autoConvert_navigator_CouchbaseClusterNodePool_To_v1alpha1_CouchbaseClusterNodePool(in *navigator.CouchbaseClusterNodePool, out *CouchbaseClusterNodePool, s conversion.Scope) error {
	out.Name = in.Name
	out.Replicas = in.Replicas
	out.Roles = *(*[]string)(unsafe.Pointer(&in.Roles))
	out.Resources = (*v1.ResourceRequirements)(unsafe.Pointer(in.Resources))
	out.Persistence = (*CouchbaseClusterPersistenceConfig)(unsafe.Pointer(in.Persistence))
	return nil
}

// Convert_navigator_CouchbaseClusterNodePool_To_v1alpha1_CouchbaseClusterNodePool is an autogenerated conversion function.
func Convert_navigator_CouchbaseClusterNodePool_To_v1alpha1_CouchbaseClusterNodePool(in *navigator.CouchbaseClusterNodePool, out *CouchbaseClusterNodePool, s conversion.Scope) error {
	return autoConvert_navigator_CouchbaseClusterNodePool_To_v1alpha1_CouchbaseClusterNodePool(in, out, s)
}

func autoConvert_v1alpha1_CouchbaseClusterPersistenceConfig_To_navigator_CouchbaseClusterPersistenceConfig(in *CouchbaseClusterPersistenceConfig, out *navigator.CouchbaseClusterPersistenceConfig, s conversion.Scope) error {
	out.Size = in.Size
	out.StorageClass = in.StorageClass
	return nil
}

// Convert_v1alpha1_CouchbaseClusterPersistenceConfig_To_navigator_CouchbaseClusterPersistenceConfig is an autogenerated conversion function.
func Convert_v1alpha1_CouchbaseClusterPersistenceConfig_To_navigator_CouchbaseClusterPersistenceConfig(in *CouchbaseClusterPersistenceConfig, out *navigator.CouchbaseClusterPersistenceConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CouchbaseClusterPersistenceConfig_To_navigator_CouchbaseClusterPersistenceConfig(in, out, s)
}

func autoConvert_navigator_CouchbaseClusterPersistenceConfig_To_v1alpha1_CouchbaseClusterPersistenceConfig(in *navigator.CouchbaseClusterPersistenceConfig, out *CouchbaseClusterPersistenceConfig, s conversion.Scope) error {
	out.Size = in.Size
	out.StorageClass = in.StorageClass
	return nil
}

// Convert_navigator_CouchbaseClusterPersistenceConfig_To_v1alpha1_CouchbaseClusterPersistenceConfig is an autogenerated conversion function.
func Convert_navigator_CouchbaseClusterPersistenceConfig_To_v1alpha1_CouchbaseClusterPersistenceConfig(in *navigator.CouchbaseClusterPersistenceConfig, out *CouchbaseClusterPersistenceConfig, s conversion.Scope) error {
	return autoConvert_navigator_CouchbaseClusterPersistenceConfig_To_v1alpha1_CouchbaseClusterPersistenceConfig(in, out, s)
}

func autoConvert_v1alpha1_CouchbaseClusterPlugin_To_navigator_CouchbaseClusterPlugin(in *CouchbaseClusterPlugin, out *navigator.CouchbaseClusterPlugin, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_CouchbaseClusterPlugin_To_navigator_CouchbaseClusterPlugin is an autogenerated conversion function.
func Convert_v1alpha1_CouchbaseClusterPlugin_To_navigator_CouchbaseClusterPlugin(in *CouchbaseClusterPlugin, out *navigator.CouchbaseClusterPlugin, s conversion.Scope) error {
	return autoConvert_v1alpha1_CouchbaseClusterPlugin_To_navigator_CouchbaseClusterPlugin(in, out, s)
}

func autoConvert_navigator_CouchbaseClusterPlugin_To_v1alpha1_CouchbaseClusterPlugin(in *navigator.CouchbaseClusterPlugin, out *CouchbaseClusterPlugin, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_navigator_CouchbaseClusterPlugin_To_v1alpha1_CouchbaseClusterPlugin is an autogenerated conversion function.
func Convert_navigator_CouchbaseClusterPlugin_To_v1alpha1_CouchbaseClusterPlugin(in *navigator.CouchbaseClusterPlugin, out *CouchbaseClusterPlugin, s conversion.Scope) error {
	return autoConvert_navigator_CouchbaseClusterPlugin_To_v1alpha1_CouchbaseClusterPlugin(in, out, s)
}

func autoConvert_v1alpha1_CouchbaseClusterSpec_To_navigator_CouchbaseClusterSpec(in *CouchbaseClusterSpec, out *navigator.CouchbaseClusterSpec, s conversion.Scope) error {
	out.Version = in.Version
	out.Plugins = *(*[]navigator.CouchbaseClusterPlugin)(unsafe.Pointer(&in.Plugins))
	out.NodePools = *(*[]navigator.CouchbaseClusterNodePool)(unsafe.Pointer(&in.NodePools))
	if err := Convert_v1alpha1_CouchbaseImage_To_navigator_CouchbaseImage(&in.Image, &out.Image, s); err != nil {
		return err
	}
	out.Sysctl = *(*[]string)(unsafe.Pointer(&in.Sysctl))
	return nil
}

// Convert_v1alpha1_CouchbaseClusterSpec_To_navigator_CouchbaseClusterSpec is an autogenerated conversion function.
func Convert_v1alpha1_CouchbaseClusterSpec_To_navigator_CouchbaseClusterSpec(in *CouchbaseClusterSpec, out *navigator.CouchbaseClusterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_CouchbaseClusterSpec_To_navigator_CouchbaseClusterSpec(in, out, s)
}

func autoConvert_navigator_CouchbaseClusterSpec_To_v1alpha1_CouchbaseClusterSpec(in *navigator.CouchbaseClusterSpec, out *CouchbaseClusterSpec, s conversion.Scope) error {
	out.Version = in.Version
	out.Plugins = *(*[]CouchbaseClusterPlugin)(unsafe.Pointer(&in.Plugins))
	out.NodePools = *(*[]CouchbaseClusterNodePool)(unsafe.Pointer(&in.NodePools))
	if err := Convert_navigator_CouchbaseImage_To_v1alpha1_CouchbaseImage(&in.Image, &out.Image, s); err != nil {
		return err
	}
	out.Sysctl = *(*[]string)(unsafe.Pointer(&in.Sysctl))
	return nil
}

// Convert_navigator_CouchbaseClusterSpec_To_v1alpha1_CouchbaseClusterSpec is an autogenerated conversion function.
func Convert_navigator_CouchbaseClusterSpec_To_v1alpha1_CouchbaseClusterSpec(in *navigator.CouchbaseClusterSpec, out *CouchbaseClusterSpec, s conversion.Scope) error {
	return autoConvert_navigator_CouchbaseClusterSpec_To_v1alpha1_CouchbaseClusterSpec(in, out, s)
}

func autoConvert_v1alpha1_CouchbaseClusterStatus_To_navigator_CouchbaseClusterStatus(in *CouchbaseClusterStatus, out *navigator.CouchbaseClusterStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_CouchbaseClusterStatus_To_navigator_CouchbaseClusterStatus is an autogenerated conversion function.
func Convert_v1alpha1_CouchbaseClusterStatus_To_navigator_CouchbaseClusterStatus(in *CouchbaseClusterStatus, out *navigator.CouchbaseClusterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CouchbaseClusterStatus_To_navigator_CouchbaseClusterStatus(in, out, s)
}

func autoConvert_navigator_CouchbaseClusterStatus_To_v1alpha1_CouchbaseClusterStatus(in *navigator.CouchbaseClusterStatus, out *CouchbaseClusterStatus, s conversion.Scope) error {
	return nil
}

// Convert_navigator_CouchbaseClusterStatus_To_v1alpha1_CouchbaseClusterStatus is an autogenerated conversion function.
func Convert_navigator_CouchbaseClusterStatus_To_v1alpha1_CouchbaseClusterStatus(in *navigator.CouchbaseClusterStatus, out *CouchbaseClusterStatus, s conversion.Scope) error {
	return autoConvert_navigator_CouchbaseClusterStatus_To_v1alpha1_CouchbaseClusterStatus(in, out, s)
}

func autoConvert_v1alpha1_CouchbaseImage_To_navigator_CouchbaseImage(in *CouchbaseImage, out *navigator.CouchbaseImage, s conversion.Scope) error {
	out.Repository = in.Repository
	out.Tag = in.Tag
	out.PullPolicy = in.PullPolicy
	out.FsGroup = in.FsGroup
	return nil
}

// Convert_v1alpha1_CouchbaseImage_To_navigator_CouchbaseImage is an autogenerated conversion function.
func Convert_v1alpha1_CouchbaseImage_To_navigator_CouchbaseImage(in *CouchbaseImage, out *navigator.CouchbaseImage, s conversion.Scope) error {
	return autoConvert_v1alpha1_CouchbaseImage_To_navigator_CouchbaseImage(in, out, s)
}

func autoConvert_navigator_CouchbaseImage_To_v1alpha1_CouchbaseImage(in *navigator.CouchbaseImage, out *CouchbaseImage, s conversion.Scope) error {
	out.Repository = in.Repository
	out.Tag = in.Tag
	out.PullPolicy = in.PullPolicy
	out.FsGroup = in.FsGroup
	return nil
}

// Convert_navigator_CouchbaseImage_To_v1alpha1_CouchbaseImage is an autogenerated conversion function.
func Convert_navigator_CouchbaseImage_To_v1alpha1_CouchbaseImage(in *navigator.CouchbaseImage, out *CouchbaseImage, s conversion.Scope) error {
	return autoConvert_navigator_CouchbaseImage_To_v1alpha1_CouchbaseImage(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster(in *ElasticsearchCluster, out *navigator.ElasticsearchCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster(in *ElasticsearchCluster, out *navigator.ElasticsearchCluster, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster(in, out, s)
}

func autoConvert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster(in *navigator.ElasticsearchCluster, out *ElasticsearchCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster is an autogenerated conversion function.
func Convert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster(in *navigator.ElasticsearchCluster, out *ElasticsearchCluster, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList(in *ElasticsearchClusterList, out *navigator.ElasticsearchClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]navigator.ElasticsearchCluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList(in *ElasticsearchClusterList, out *navigator.ElasticsearchClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList(in *navigator.ElasticsearchClusterList, out *ElasticsearchClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ElasticsearchCluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList(in *navigator.ElasticsearchClusterList, out *ElasticsearchClusterList, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool(in *ElasticsearchClusterNodePool, out *navigator.ElasticsearchClusterNodePool, s conversion.Scope) error {
	out.Name = in.Name
	out.Replicas = in.Replicas
	out.Roles = *(*[]string)(unsafe.Pointer(&in.Roles))
	out.Resources = (*v1.ResourceRequirements)(unsafe.Pointer(in.Resources))
	out.Persistence = (*navigator.ElasticsearchClusterPersistenceConfig)(unsafe.Pointer(in.Persistence))
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool(in *ElasticsearchClusterNodePool, out *navigator.ElasticsearchClusterNodePool, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool(in *navigator.ElasticsearchClusterNodePool, out *ElasticsearchClusterNodePool, s conversion.Scope) error {
	out.Name = in.Name
	out.Replicas = in.Replicas
	out.Roles = *(*[]string)(unsafe.Pointer(&in.Roles))
	out.Resources = (*v1.ResourceRequirements)(unsafe.Pointer(in.Resources))
	out.Persistence = (*ElasticsearchClusterPersistenceConfig)(unsafe.Pointer(in.Persistence))
	return nil
}

// Convert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool(in *navigator.ElasticsearchClusterNodePool, out *ElasticsearchClusterNodePool, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig(in *ElasticsearchClusterPersistenceConfig, out *navigator.ElasticsearchClusterPersistenceConfig, s conversion.Scope) error {
	out.Size = in.Size
	out.StorageClass = in.StorageClass
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig(in *ElasticsearchClusterPersistenceConfig, out *navigator.ElasticsearchClusterPersistenceConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig(in *navigator.ElasticsearchClusterPersistenceConfig, out *ElasticsearchClusterPersistenceConfig, s conversion.Scope) error {
	out.Size = in.Size
	out.StorageClass = in.StorageClass
	return nil
}

// Convert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig(in *navigator.ElasticsearchClusterPersistenceConfig, out *ElasticsearchClusterPersistenceConfig, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterPlugin_To_navigator_ElasticsearchClusterPlugin(in *ElasticsearchClusterPlugin, out *navigator.ElasticsearchClusterPlugin, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterPlugin_To_navigator_ElasticsearchClusterPlugin is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterPlugin_To_navigator_ElasticsearchClusterPlugin(in *ElasticsearchClusterPlugin, out *navigator.ElasticsearchClusterPlugin, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterPlugin_To_navigator_ElasticsearchClusterPlugin(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterPlugin_To_v1alpha1_ElasticsearchClusterPlugin(in *navigator.ElasticsearchClusterPlugin, out *ElasticsearchClusterPlugin, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_navigator_ElasticsearchClusterPlugin_To_v1alpha1_ElasticsearchClusterPlugin is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterPlugin_To_v1alpha1_ElasticsearchClusterPlugin(in *navigator.ElasticsearchClusterPlugin, out *ElasticsearchClusterPlugin, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterPlugin_To_v1alpha1_ElasticsearchClusterPlugin(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec(in *ElasticsearchClusterSpec, out *navigator.ElasticsearchClusterSpec, s conversion.Scope) error {
	out.Version = in.Version
	out.Plugins = *(*[]navigator.ElasticsearchClusterPlugin)(unsafe.Pointer(&in.Plugins))
	out.NodePools = *(*[]navigator.ElasticsearchClusterNodePool)(unsafe.Pointer(&in.NodePools))
	if err := Convert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage(&in.Image, &out.Image, s); err != nil {
		return err
	}
	out.Sysctl = *(*[]string)(unsafe.Pointer(&in.Sysctl))
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec(in *ElasticsearchClusterSpec, out *navigator.ElasticsearchClusterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec(in *navigator.ElasticsearchClusterSpec, out *ElasticsearchClusterSpec, s conversion.Scope) error {
	out.Version = in.Version
	out.Plugins = *(*[]ElasticsearchClusterPlugin)(unsafe.Pointer(&in.Plugins))
	out.NodePools = *(*[]ElasticsearchClusterNodePool)(unsafe.Pointer(&in.NodePools))
	if err := Convert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage(&in.Image, &out.Image, s); err != nil {
		return err
	}
	out.Sysctl = *(*[]string)(unsafe.Pointer(&in.Sysctl))
	return nil
}

// Convert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec(in *navigator.ElasticsearchClusterSpec, out *ElasticsearchClusterSpec, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus(in *ElasticsearchClusterStatus, out *navigator.ElasticsearchClusterStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus(in *ElasticsearchClusterStatus, out *navigator.ElasticsearchClusterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus(in *navigator.ElasticsearchClusterStatus, out *ElasticsearchClusterStatus, s conversion.Scope) error {
	return nil
}

// Convert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus(in *navigator.ElasticsearchClusterStatus, out *ElasticsearchClusterStatus, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage(in *ElasticsearchImage, out *navigator.ElasticsearchImage, s conversion.Scope) error {
	out.Repository = in.Repository
	out.Tag = in.Tag
	out.PullPolicy = in.PullPolicy
	out.FsGroup = in.FsGroup
	return nil
}

// Convert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage(in *ElasticsearchImage, out *navigator.ElasticsearchImage, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage(in, out, s)
}

func autoConvert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage(in *navigator.ElasticsearchImage, out *ElasticsearchImage, s conversion.Scope) error {
	out.Repository = in.Repository
	out.Tag = in.Tag
	out.PullPolicy = in.PullPolicy
	out.FsGroup = in.FsGroup
	return nil
}

// Convert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage is an autogenerated conversion function.
func Convert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage(in *navigator.ElasticsearchImage, out *ElasticsearchImage, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage(in, out, s)
}
