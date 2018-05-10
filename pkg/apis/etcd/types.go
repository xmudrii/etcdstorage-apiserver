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

package etcd

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EtcdStorageList is a list of EtcdStorage objects.
type EtcdStorageList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []EtcdStorage
}

// ReferenceType is a string describing reference.
type ReferenceType string

const (
	// EtcdStorageReferenceType is reference type of etcdStorage.
	EtcdStorageReferenceType = ReferenceType("EtcdStorage")
)

// EtcdStorageSpec is etcdStorage's spec.
type EtcdStorageSpec struct {
	// A name of another etcd storage.
	EtcdStorageReference string
	// The reference type.
	ReferenceType ReferenceType
}

// EtcdStorageStatus is current status of etcdStorage object.
type EtcdStorageStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EtcdStorage describe an etcdStorage object.
type EtcdStorage struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   EtcdStorageSpec
	Status EtcdStorageStatus
}
