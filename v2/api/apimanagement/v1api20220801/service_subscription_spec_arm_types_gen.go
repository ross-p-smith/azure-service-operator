// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Service_Subscription_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Subscription contract properties.
	Properties *SubscriptionCreateParameterProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Service_Subscription_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (subscription Service_Subscription_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (subscription *Service_Subscription_Spec_ARM) GetName() string {
	return subscription.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/subscriptions"
func (subscription *Service_Subscription_Spec_ARM) GetType() string {
	return "Microsoft.ApiManagement/service/subscriptions"
}

// Parameters supplied to the Create subscription operation.
type SubscriptionCreateParameterProperties_ARM struct {
	// AllowTracing: Determines whether tracing can be enabled
	AllowTracing *bool `json:"allowTracing,omitempty"`

	// DisplayName: Subscription name.
	DisplayName *string `json:"displayName,omitempty"`

	// OwnerId: User (user id path) for whom subscription is being created in form /users/{userId}
	OwnerId *string `json:"ownerId,omitempty"`

	// PrimaryKey: Primary subscription key. If not specified during request key will be generated automatically.
	PrimaryKey *string `json:"primaryKey,omitempty"`

	// Scope: Scope like /products/{productId} or /apis or /apis/{apiId}.
	Scope *string `json:"scope,omitempty"`

	// SecondaryKey: Secondary subscription key. If not specified during request key will be generated automatically.
	SecondaryKey *string `json:"secondaryKey,omitempty"`

	// State: Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible
	// states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber
	// cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has
	// not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, *
	// cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription
	// reached its expiration date and was deactivated.
	State *SubscriptionCreateParameterProperties_State `json:"state,omitempty"`
}

// +kubebuilder:validation:Enum={"active","cancelled","expired","rejected","submitted","suspended"}
type SubscriptionCreateParameterProperties_State string

const (
	SubscriptionCreateParameterProperties_State_Active    = SubscriptionCreateParameterProperties_State("active")
	SubscriptionCreateParameterProperties_State_Cancelled = SubscriptionCreateParameterProperties_State("cancelled")
	SubscriptionCreateParameterProperties_State_Expired   = SubscriptionCreateParameterProperties_State("expired")
	SubscriptionCreateParameterProperties_State_Rejected  = SubscriptionCreateParameterProperties_State("rejected")
	SubscriptionCreateParameterProperties_State_Submitted = SubscriptionCreateParameterProperties_State("submitted")
	SubscriptionCreateParameterProperties_State_Suspended = SubscriptionCreateParameterProperties_State("suspended")
)
