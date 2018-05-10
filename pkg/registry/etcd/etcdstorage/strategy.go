/*
Copyright 2017 The Kubernetes Authors.

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

package etcdstorage

import (
	"context"
	"fmt"

	"github.com/xmudrii/etcdstorage-apiserver/pkg/apis/etcd/validation"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	"github.com/xmudrii/etcdstorage-apiserver/pkg/apis/etcd"
)

// NewStrategy creates and returns a etcdStorageStrategy instance
func NewStrategy(typer runtime.ObjectTyper) etcdStorageStrategy {
	return etcdStorageStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a EtcdStorage
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	apiserver, ok := obj.(*etcd.EtcdStorage)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a EtcdStorage")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), apiserver.Initializers != nil, nil
}

// MatchEtcdStorage is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchEtcdStorage(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *etcd.EtcdStorage) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type etcdStorageStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (etcdStorageStrategy) NamespaceScoped() bool {
	return true
}

func (etcdStorageStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (etcdStorageStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (etcdStorageStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	etcdStorage := obj.(*etcd.EtcdStorage)
	return validation.ValidateEtcdStorage(etcdStorage)
}

func (etcdStorageStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (etcdStorageStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (etcdStorageStrategy) Canonicalize(obj runtime.Object) {
}

func (etcdStorageStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
