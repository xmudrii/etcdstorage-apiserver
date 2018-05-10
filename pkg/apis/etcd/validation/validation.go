/*
Copyright 2016 The Kubernetes Authors.

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

package validation

import (
	"github.com/xmudrii/etcdstorage-apiserver/pkg/apis/etcd"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateEtcdStorage validates a EtcdStorage.
func ValidateEtcdStorage(f *etcd.EtcdStorage) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, ValidateEtcdStorageSpec(&f.Spec, field.NewPath("spec"))...)

	return allErrs
}

// ValidateEtcdStorageSpec validates a EtcdStorageSpec.
func ValidateEtcdStorageSpec(s *etcd.EtcdStorageSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(s.EtcdStorageReference) != 0 && s.ReferenceType != etcd.EtcdStorageReferenceType {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("etcdStorageReference"), s.EtcdStorageReference, "cannot be set if referenceType is not EtcdStorage"))
	} else if len(s.EtcdStorageReference) == 0 && s.ReferenceType == etcd.EtcdStorageReferenceType {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("etcdStorageReference"), s.EtcdStorageReference, "cannot be empty if referenceType is EtcdStorage"))
	}

	if len(s.ReferenceType) != 0 && s.ReferenceType != etcd.EtcdStorageReferenceType {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("referenceType"), s.ReferenceType, "must be EtcdStorage"))
	}

	return allErrs
}
