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
	"fmt"
	"net"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var iprangelog = logf.Log.WithName("iprange-resource")

func (r *IPRange) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update;delete,path=/validate-allocator-k8s-io-v1-iprange,mutating=false,failurePolicy=fail,groups=allocator.k8s.io,resources=ipranges,versions=v1,name=viprange.kb.io

var _ webhook.Validator = &IPRange{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *IPRange) ValidateCreate() error {
	iprangelog.Info("validate create", "name", r.Name)
	_, _, err := net.ParseCIDR(r.Spec.Range)
	return err
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *IPRange) ValidateUpdate(old runtime.Object) error {
	iprangelog.Info("validate update", "name", r.Name)
	return fmt.Errorf("IPRange can not be changed after creation")
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *IPRange) ValidateDelete() error {
	iprangelog.Info("validate delete", "name", r.Name)
	// TODO: IPRange only can be deleted if there are NO IPAddress referencing it
	return fmt.Errorf("IPRange can not be deleted")
}
