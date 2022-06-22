/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RSCTSpec defines the desired state of RSCT
type RSCTSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RSCT. Edit rsct_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// Represents an output from rmcdomainstatus
// See https://www.ibm.com/docs/en/aix/7.2?topic=r-rmcdomainstatus-command
type NodePodStatus struct {
	TokenOne           string `json:"nodePodToken1,omitempty"`
	TokenTwo           string `json:"nodePodToken2,omitempty"`
	NodeID             string `json:"nodeID,omitempty"`
	InternalNodeNumber string `json:"internalNodeNumber,omitempty"`
	IPAddress          string `json:"ipAddress,omitempty"`
}

type NodePod struct {
	PodName string        `json:"nodePodName"`
	Status  NodePodStatus `json:"nodePodStatus,omitempty"`
}

// RSCTStatus defines the observed state of RSCT
type RSCTStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	NodePods []NodePod `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RSCT is the Schema for the rscts API
type RSCT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RSCTSpec   `json:"spec,omitempty"`
	Status RSCTStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RSCTList contains a list of RSCT
type RSCTList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RSCT `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RSCT{}, &RSCTList{})
}
