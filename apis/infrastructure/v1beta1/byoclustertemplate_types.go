// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// BYOClusterTemplateSpec defines the desired state of BYOClusterTemplate.
type BYOClusterTemplateSpec struct {
	Template BYOClusterTemplateResource `json:"template"`
}

//+kubebuilder:object:root=true

// BYOClusterTemplate is the Schema for the byoclustertemplates API.
type BYOClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec BYOClusterTemplateSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// BYOClusterTemplateList contains a list of BYOClusterTemplate.
type BYOClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BYOClusterTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BYOClusterTemplate{}, &BYOClusterTemplateList{})
}

// BYOClusterTemplateResource describes the data needed to create a BYOCluster from a template.
type BYOClusterTemplateResource struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	ObjectMeta clusterv1.ObjectMeta `json:"metadata,omitempty"`
	Spec       ByoClusterSpec       `json:"spec"`
}
