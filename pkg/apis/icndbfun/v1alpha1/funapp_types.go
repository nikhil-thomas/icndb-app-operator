package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FunAppSpec defines the desired state of FunApp
// +k8s:openapi-gen=true
type FunAppSpec struct {
	// Funpods specify number of replicas in the deployment created
	Funpods int32 `json:"funpods"`
}

// FunAppStatus defines the observed state of FunApp
// +k8s:openapi-gen=true
type FunAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FunApp is the Schema for the funapps API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type FunApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FunAppSpec   `json:"spec,omitempty"`
	Status FunAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FunAppList contains a list of FunApp
type FunAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FunApp{}, &FunAppList{})
}
