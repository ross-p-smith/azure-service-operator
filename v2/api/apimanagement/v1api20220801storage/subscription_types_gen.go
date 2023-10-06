// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=apimanagement.azure.com,resources=subscriptions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apimanagement.azure.com,resources={subscriptions/status,subscriptions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220801.Subscription
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimsubscriptions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Service_Subscription_Spec   `json:"spec,omitempty"`
	Status            Service_Subscription_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Subscription{}

// GetConditions returns the conditions of the resource
func (subscription *Subscription) GetConditions() conditions.Conditions {
	return subscription.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (subscription *Subscription) SetConditions(conditions conditions.Conditions) {
	subscription.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Subscription{}

// AzureName returns the Azure name of the resource
func (subscription *Subscription) AzureName() string {
	return subscription.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (subscription Subscription) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (subscription *Subscription) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (subscription *Subscription) GetSpec() genruntime.ConvertibleSpec {
	return &subscription.Spec
}

// GetStatus returns the status of this resource
func (subscription *Subscription) GetStatus() genruntime.ConvertibleStatus {
	return &subscription.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/subscriptions"
func (subscription *Subscription) GetType() string {
	return "Microsoft.ApiManagement/service/subscriptions"
}

// NewEmptyStatus returns a new empty (blank) status
func (subscription *Subscription) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Service_Subscription_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (subscription *Subscription) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(subscription.Spec)
	return subscription.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (subscription *Subscription) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Service_Subscription_STATUS); ok {
		subscription.Status = *st
		return nil
	}

	// Convert status to required version
	var st Service_Subscription_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	subscription.Status = st
	return nil
}

// Hub marks that this Subscription is the hub type for conversion
func (subscription *Subscription) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (subscription *Subscription) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: subscription.Spec.OriginalVersion,
		Kind:    "Subscription",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220801.Subscription
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimsubscriptions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Storage version of v1api20220801.Service_Subscription_Spec
type Service_Subscription_Spec struct {
	AllowTracing *bool `json:"allowTracing,omitempty"`

	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:Pattern="^[^*#&+:<>?]+$"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                    `json:"azureName,omitempty"`
	DisplayName     *string                   `json:"displayName,omitempty"`
	OperatorSpec    *SubscriptionOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                    `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner                  *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	OwnerId                *string                            `json:"ownerId,omitempty"`
	PrimaryKey             *string                            `json:"primaryKey,omitempty" optionalConfigMapPair:"PrimaryKey"`
	PrimaryKeyFromConfig   *genruntime.ConfigMapReference     `json:"primaryKeyFromConfig,omitempty" optionalConfigMapPair:"PrimaryKey"`
	PropertyBag            genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Scope                  *string                            `json:"scope,omitempty"`
	SecondaryKey           *string                            `json:"secondaryKey,omitempty" optionalConfigMapPair:"SecondaryKey"`
	SecondaryKeyFromConfig *genruntime.ConfigMapReference     `json:"secondaryKeyFromConfig,omitempty" optionalConfigMapPair:"SecondaryKey"`
	State                  *string                            `json:"state,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Service_Subscription_Spec{}

// ConvertSpecFrom populates our Service_Subscription_Spec from the provided source
func (subscription *Service_Subscription_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(subscription)
}

// ConvertSpecTo populates the provided destination from our Service_Subscription_Spec
func (subscription *Service_Subscription_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(subscription)
}

// Storage version of v1api20220801.Service_Subscription_STATUS
type Service_Subscription_STATUS struct {
	AllowTracing     *bool                  `json:"allowTracing,omitempty"`
	Conditions       []conditions.Condition `json:"conditions,omitempty"`
	CreatedDate      *string                `json:"createdDate,omitempty"`
	DisplayName      *string                `json:"displayName,omitempty"`
	EndDate          *string                `json:"endDate,omitempty"`
	ExpirationDate   *string                `json:"expirationDate,omitempty"`
	Id               *string                `json:"id,omitempty"`
	Name             *string                `json:"name,omitempty"`
	NotificationDate *string                `json:"notificationDate,omitempty"`
	OwnerId          *string                `json:"ownerId,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Scope            *string                `json:"scope,omitempty"`
	StartDate        *string                `json:"startDate,omitempty"`
	State            *string                `json:"state,omitempty"`
	StateComment     *string                `json:"stateComment,omitempty"`
	Type             *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Service_Subscription_STATUS{}

// ConvertStatusFrom populates our Service_Subscription_STATUS from the provided source
func (subscription *Service_Subscription_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(subscription)
}

// ConvertStatusTo populates the provided destination from our Service_Subscription_STATUS
func (subscription *Service_Subscription_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(subscription)
}

// Storage version of v1api20220801.SubscriptionOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SubscriptionOperatorSpec struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Secrets     *SubscriptionOperatorSecrets `json:"secrets,omitempty"`
}

// Storage version of v1api20220801.SubscriptionOperatorSecrets
type SubscriptionOperatorSecrets struct {
	PrimaryKey   *genruntime.SecretDestination `json:"primaryKey,omitempty"`
	PropertyBag  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecondaryKey *genruntime.SecretDestination `json:"secondaryKey,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}
