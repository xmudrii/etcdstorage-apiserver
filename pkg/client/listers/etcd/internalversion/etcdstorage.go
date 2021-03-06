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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	etcd "github.com/xmudrii/etcdstorage-apiserver/pkg/apis/etcd"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EtcdStorageLister helps list EtcdStorages.
type EtcdStorageLister interface {
	// List lists all EtcdStorages in the indexer.
	List(selector labels.Selector) (ret []*etcd.EtcdStorage, err error)
	// EtcdStorages returns an object that can list and get EtcdStorages.
	EtcdStorages(namespace string) EtcdStorageNamespaceLister
	EtcdStorageListerExpansion
}

// etcdStorageLister implements the EtcdStorageLister interface.
type etcdStorageLister struct {
	indexer cache.Indexer
}

// NewEtcdStorageLister returns a new EtcdStorageLister.
func NewEtcdStorageLister(indexer cache.Indexer) EtcdStorageLister {
	return &etcdStorageLister{indexer: indexer}
}

// List lists all EtcdStorages in the indexer.
func (s *etcdStorageLister) List(selector labels.Selector) (ret []*etcd.EtcdStorage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*etcd.EtcdStorage))
	})
	return ret, err
}

// EtcdStorages returns an object that can list and get EtcdStorages.
func (s *etcdStorageLister) EtcdStorages(namespace string) EtcdStorageNamespaceLister {
	return etcdStorageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EtcdStorageNamespaceLister helps list and get EtcdStorages.
type EtcdStorageNamespaceLister interface {
	// List lists all EtcdStorages in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*etcd.EtcdStorage, err error)
	// Get retrieves the EtcdStorage from the indexer for a given namespace and name.
	Get(name string) (*etcd.EtcdStorage, error)
	EtcdStorageNamespaceListerExpansion
}

// etcdStorageNamespaceLister implements the EtcdStorageNamespaceLister
// interface.
type etcdStorageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EtcdStorages in the indexer for a given namespace.
func (s etcdStorageNamespaceLister) List(selector labels.Selector) (ret []*etcd.EtcdStorage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*etcd.EtcdStorage))
	})
	return ret, err
}

// Get retrieves the EtcdStorage from the indexer for a given namespace and name.
func (s etcdStorageNamespaceLister) Get(name string) (*etcd.EtcdStorage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(etcd.Resource("etcdstorage"), name)
	}
	return obj.(*etcd.EtcdStorage), nil
}
