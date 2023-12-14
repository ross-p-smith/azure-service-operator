// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Service_User_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: User entity create contract properties.
	Properties *UserCreateParameterProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Service_User_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (user Service_User_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (user *Service_User_Spec_ARM) GetName() string {
	return user.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/users"
func (user *Service_User_Spec_ARM) GetType() string {
	return "Microsoft.ApiManagement/service/users"
}

// Parameters supplied to the Create User operation.
type UserCreateParameterProperties_ARM struct {
	// AppType: Determines the type of application which send the create user request. Default is legacy portal.
	AppType *UserCreateParameterProperties_AppType `json:"appType,omitempty"`

	// Confirmation: Determines the type of confirmation e-mail that will be sent to the newly created user.
	Confirmation *UserCreateParameterProperties_Confirmation `json:"confirmation,omitempty"`

	// Email: Email address. Must not be empty and must be unique within the service instance.
	Email *string `json:"email,omitempty"`

	// FirstName: First name.
	FirstName *string `json:"firstName,omitempty"`

	// Identities: Collection of user identities.
	Identities []UserIdentityContract_ARM `json:"identities,omitempty"`

	// LastName: Last name.
	LastName *string `json:"lastName,omitempty"`

	// Note: Optional note about a user set by the administrator.
	Note *string `json:"note,omitempty"`

	// Password: User Password. If no value is provided, a default password is generated.
	Password *string `json:"password,omitempty"`

	// State: Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer
	// portal or call any APIs of subscribed products. Default state is Active.
	State *UserCreateParameterProperties_State `json:"state,omitempty"`
}

// +kubebuilder:validation:Enum={"developerPortal","portal"}
type UserCreateParameterProperties_AppType string

const (
	UserCreateParameterProperties_AppType_DeveloperPortal = UserCreateParameterProperties_AppType("developerPortal")
	UserCreateParameterProperties_AppType_Portal          = UserCreateParameterProperties_AppType("portal")
)

// +kubebuilder:validation:Enum={"invite","signup"}
type UserCreateParameterProperties_Confirmation string

const (
	UserCreateParameterProperties_Confirmation_Invite = UserCreateParameterProperties_Confirmation("invite")
	UserCreateParameterProperties_Confirmation_Signup = UserCreateParameterProperties_Confirmation("signup")
)

// +kubebuilder:validation:Enum={"active","blocked","deleted","pending"}
type UserCreateParameterProperties_State string

const (
	UserCreateParameterProperties_State_Active  = UserCreateParameterProperties_State("active")
	UserCreateParameterProperties_State_Blocked = UserCreateParameterProperties_State("blocked")
	UserCreateParameterProperties_State_Deleted = UserCreateParameterProperties_State("deleted")
	UserCreateParameterProperties_State_Pending = UserCreateParameterProperties_State("pending")
)

// User identity details.
type UserIdentityContract_ARM struct {
	Id *string `json:"id,omitempty"`

	// Provider: Identity provider name.
	Provider *string `json:"provider,omitempty"`
}