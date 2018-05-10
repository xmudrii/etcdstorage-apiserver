/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	etcd_v1alpha1 "github.com/xmudrii/etcdstorage-apiserver/pkg/apis/etcd/v1alpha1"
	versioned "github.com/xmudrii/etcdstorage-apiserver/pkg/client/clientset/versioned"
	internalinterfaces "github.com/xmudrii/etcdstorage-apiserver/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/xmudrii/etcdstorage-apiserver/pkg/client/listers/etcd/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EtcdStorageInformer provides access to a shared informer and lister for
// EtcdStorages.
type EtcdStorageInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.EtcdStorageLister
}

type etcdStorageInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEtcdStorageInformer constructs a new informer for EtcdStorage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEtcdStorageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEtcdStorageInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEtcdStorageInformer constructs a new informer for EtcdStorage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEtcdStorageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EtcdV1alpha1().EtcdStorages(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EtcdV1alpha1().EtcdStorages(namespace).Watch(options)
			},
		},
		&etcd_v1alpha1.EtcdStorage{},
		resyncPeriod,
		indexers,
	)
}

func (f *etcdStorageInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEtcdStorageInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *etcdStorageInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&etcd_v1alpha1.EtcdStorage{}, f.defaultInformer)
}

func (f *etcdStorageInformer) Lister() v1alpha1.EtcdStorageLister {
	return v1alpha1.NewEtcdStorageLister(f.Informer().GetIndexer())
}
