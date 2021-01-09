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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IPAddressSpec defines the desired state of IPAddress
type IPAddressSpec struct {
	// IPRangeName references the IPRange that this address belongs too
	IPRangeName string `json:"ipRangeName,omitempty"`
}

// IPAddressStatus defines the observed state of IPAddress
type IPAddressStatus struct {
	State IPAddressState
}

// IPAddressState defines the state of the IP address
// +kubebuilder:validation:Enum=Pending;Allocated;Free
type IPAddressState string

// These are the valid statuses of IPAddresses.
const (
	// IPPending means the IP has been allocated by the system but not allocated to a Service yet.
	IPPending IPAddressState = "Pending"
	// IPAllocated means the IP belongs to a Service and has been persisted.
	IPAllocated IPAddressState = "Allocated"
	// IPFree means that IP is not being used.
	IPFree IPAddressState = "Free"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status

// IPAddress is the Schema for the ipaddresses API
type IPAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IPAddressSpec   `json:"spec,omitempty"`
	Status IPAddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPAddressList contains a list of IPAddress
type IPAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPAddress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IPAddress{}, &IPAddressList{})
}
