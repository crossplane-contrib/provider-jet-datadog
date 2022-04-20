/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type JSONObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type JSONParameters struct {

	// The JSON formatted definition of the monitor.
	// +kubebuilder:validation:Required
	Monitor *string `json:"monitor" tf:"monitor,omitempty"`

	// The URL of the monitor.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// JSONSpec defines the desired state of JSON
type JSONSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JSONParameters `json:"forProvider"`
}

// JSONStatus defines the observed state of JSON.
type JSONStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JSONObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JSON is the Schema for the JSONs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type JSON struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JSONSpec   `json:"spec"`
	Status            JSONStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JSONList contains a list of JSONs
type JSONList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JSON `json:"items"`
}

// Repository type metadata.
var (
	JSON_Kind             = "JSON"
	JSON_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JSON_Kind}.String()
	JSON_KindAPIVersion   = JSON_Kind + "." + CRDGroupVersion.String()
	JSON_GroupVersionKind = CRDGroupVersion.WithKind(JSON_Kind)
)

func init() {
	SchemeBuilder.Register(&JSON{}, &JSONList{})
}