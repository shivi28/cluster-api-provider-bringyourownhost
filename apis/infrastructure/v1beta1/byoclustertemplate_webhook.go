// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)


func (byoClusterTemplate *BYOClusterTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(byoClusterTemplate).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-infrastructure-cluster-x-k8s-io-v1beta1-byocluster,mutating=true,failurePolicy=fail,sideEffects=None,groups=infrastructure.cluster.x-k8s.io,resources=byoclusters,verbs=create;update,versions=v1beta1,name=mbyocluster.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &BYOClusterTemplate{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (byoClusterTemplate *BYOClusterTemplate) Default() {
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-infrastructure-cluster-x-k8s-io-v1beta1-byocluster,mutating=false,failurePolicy=fail,sideEffects=None,groups=infrastructure.cluster.x-k8s.io,resources=byoclusters,verbs=create;update,versions=v1beta1,name=vbyocluster.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &BYOClusterTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (byoClusterTemplate *BYOClusterTemplate) ValidateCreate() error {
	byoclusterlog.Info("validate create", "name", byoClusterTemplate.Name)
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (byoClusterTemplate *BYOClusterTemplate) ValidateUpdate(old runtime.Object) error {
	byoclusterlog.Info("validate update", "name", byoClusterTemplate.Name)
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (byoClusterTemplate *BYOClusterTemplate) ValidateDelete() error {
	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
