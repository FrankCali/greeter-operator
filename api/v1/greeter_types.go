/*
Copyright 2025.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GreeterSpec defines the desired state of Greeter.
type GreeterSpec struct {
	// Name is the person's name to greet
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:Pattern=^[A-Za-z\s]+$
	Name string `json:"name"`

	// Message is an optional custom message
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MaxLength=100
	Message string `json:"message,omitempty"`
}

// GreeterStatus defines the observed state of Greeter.
type GreeterStatus struct {
	// Greeting contains the generated greeting message
	Greeting string `json:"greeting,omitempty"`

	// Ready indicates whether the ConfigMap has been created
	Ready bool `json:"ready,omitempty"`

	// LastUpdated timestamp of the last reconciliation
	LastUpdated string `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Name",type="string",JSONPath=".spec.name"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".spec.message"
// +kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.ready"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Greeter is the Schema for the greeters API.
type Greeter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GreeterSpec   `json:"spec,omitempty"`
	Status GreeterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GreeterList contains a list of Greeter.
type GreeterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Greeter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Greeter{}, &GreeterList{})
}
