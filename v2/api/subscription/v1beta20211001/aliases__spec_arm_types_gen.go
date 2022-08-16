// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Aliases_SpecARM struct {
	// Name: AliasName is the name for the subscription creation request. Note that this is not the same as subscription name
	// and this doesn’t have any other lifecycle need beyond the request for subscription creation.
	Name string `json:"name,omitempty"`

	// Properties: Put subscription properties.
	Properties *PutAliasRequestPropertiesARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Aliases_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (aliases Aliases_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (aliases *Aliases_SpecARM) GetName() string {
	return aliases.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Subscription/aliases"
func (aliases *Aliases_SpecARM) GetType() string {
	return "Microsoft.Subscription/aliases"
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.Subscription.json#/definitions/PutAliasRequestProperties
type PutAliasRequestPropertiesARM struct {
	// AdditionalProperties: Put subscription additional properties.
	AdditionalProperties *PutAliasRequestAdditionalPropertiesARM `json:"additionalProperties,omitempty"`

	// BillingScope: Billing scope of the subscription.
	// For CustomerLed and FieldLed -
	// /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
	// For PartnerLed - /billingAccounts/{billingAccountName}/customers/{customerName}
	// For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
	BillingScope *string `json:"billingScope,omitempty"`

	// DisplayName: The friendly name of the subscription.
	DisplayName *string `json:"displayName,omitempty"`

	// ResellerId: Reseller Id
	ResellerId *string `json:"resellerId,omitempty"`

	// SubscriptionId: This parameter can be used to create alias for existing subscription Id
	SubscriptionId *string                            `json:"subscriptionId,omitempty"`
	Workload       *PutAliasRequestPropertiesWorkload `json:"workload,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.Subscription.json#/definitions/PutAliasRequestAdditionalProperties
type PutAliasRequestAdditionalPropertiesARM struct {
	// ManagementGroupId: Management group Id for the subscription.
	ManagementGroupId *string `json:"managementGroupId,omitempty"`

	// SubscriptionOwnerId: Owner Id of the subscription
	SubscriptionOwnerId *string `json:"subscriptionOwnerId,omitempty"`

	// SubscriptionTenantId: Tenant Id of the subscription
	SubscriptionTenantId *string `json:"subscriptionTenantId,omitempty"`

	// Tags: Tags for the subscription
	Tags map[string]string `json:"tags,omitempty"`
}

// +kubebuilder:validation:Enum={"DevTest","Production"}
type PutAliasRequestPropertiesWorkload string

const (
	PutAliasRequestPropertiesWorkloadDevTest    = PutAliasRequestPropertiesWorkload("DevTest")
	PutAliasRequestPropertiesWorkloadProduction = PutAliasRequestPropertiesWorkload("Production")
)