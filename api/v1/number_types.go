/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
)

var _ resource.Object = &Number{}
var _ resource.ObjectList = &NumberList{}

// GetObjectMeta returns the object meta reference.
func (n *Number) GetObjectMeta() *metav1.ObjectMeta {
	return &n.ObjectMeta
}
func (n *Number) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "rnd.glrf.xyz",
		Version:  "v1",
		Resource: "numbers",
	}
}
func (n *Number) IsStorageVersion() bool {
	return true
}
func (n *Number) NamespaceScoped() bool {
	return true
}
func (n *Number) New() runtime.Object {
	return &Number{}
}
func (n *Number) NewList() runtime.Object {
	return &NumberList{}
}

func (in *NumberList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// NumberSpec defines the desired state of Number
type NumberSpec struct {
	Int int `json:"int,omitempty"`
}

// NumberStatus defines the observed state of Number
type NumberStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Number is the Schema for the numbers API
type Number struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NumberSpec   `json:"spec,omitempty"`
	Status NumberStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NumberList contains a list of Number
type NumberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Number `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Number{}, &NumberList{})
}
