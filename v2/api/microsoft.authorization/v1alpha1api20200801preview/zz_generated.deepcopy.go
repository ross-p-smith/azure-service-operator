//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20200801preview

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignment) DeepCopyInto(out *RoleAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignment.
func (in *RoleAssignment) DeepCopy() *RoleAssignment {
	if in == nil {
		return nil
	}
	out := new(RoleAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentList) DeepCopyInto(out *RoleAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentList.
func (in *RoleAssignmentList) DeepCopy() *RoleAssignmentList {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentPropertiesARM) DeepCopyInto(out *RoleAssignmentPropertiesARM) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.ConditionVersion != nil {
		in, out := &in.ConditionVersion, &out.ConditionVersion
		*out = new(string)
		**out = **in
	}
	if in.DelegatedManagedIdentityResourceId != nil {
		in, out := &in.DelegatedManagedIdentityResourceId, &out.DelegatedManagedIdentityResourceId
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(RoleAssignmentPropertiesPrincipalType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentPropertiesARM.
func (in *RoleAssignmentPropertiesARM) DeepCopy() *RoleAssignmentPropertiesARM {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentPropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentProperties_StatusARM) DeepCopyInto(out *RoleAssignmentProperties_StatusARM) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.ConditionVersion != nil {
		in, out := &in.ConditionVersion, &out.ConditionVersion
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedOn != nil {
		in, out := &in.CreatedOn, &out.CreatedOn
		*out = new(string)
		**out = **in
	}
	if in.DelegatedManagedIdentityResourceId != nil {
		in, out := &in.DelegatedManagedIdentityResourceId, &out.DelegatedManagedIdentityResourceId
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(RoleAssignmentPropertiesStatusPrincipalType)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.UpdatedBy != nil {
		in, out := &in.UpdatedBy, &out.UpdatedBy
		*out = new(string)
		**out = **in
	}
	if in.UpdatedOn != nil {
		in, out := &in.UpdatedOn, &out.UpdatedOn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentProperties_StatusARM.
func (in *RoleAssignmentProperties_StatusARM) DeepCopy() *RoleAssignmentProperties_StatusARM {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentProperties_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignment_Status) DeepCopyInto(out *RoleAssignment_Status) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.ConditionVersion != nil {
		in, out := &in.ConditionVersion, &out.ConditionVersion
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedOn != nil {
		in, out := &in.CreatedOn, &out.CreatedOn
		*out = new(string)
		**out = **in
	}
	if in.DelegatedManagedIdentityResourceId != nil {
		in, out := &in.DelegatedManagedIdentityResourceId, &out.DelegatedManagedIdentityResourceId
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(RoleAssignmentPropertiesStatusPrincipalType)
		**out = **in
	}
	if in.RoleDefinitionId != nil {
		in, out := &in.RoleDefinitionId, &out.RoleDefinitionId
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UpdatedBy != nil {
		in, out := &in.UpdatedBy, &out.UpdatedBy
		*out = new(string)
		**out = **in
	}
	if in.UpdatedOn != nil {
		in, out := &in.UpdatedOn, &out.UpdatedOn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignment_Status.
func (in *RoleAssignment_Status) DeepCopy() *RoleAssignment_Status {
	if in == nil {
		return nil
	}
	out := new(RoleAssignment_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignment_StatusARM) DeepCopyInto(out *RoleAssignment_StatusARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(RoleAssignmentProperties_StatusARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignment_StatusARM.
func (in *RoleAssignment_StatusARM) DeepCopy() *RoleAssignment_StatusARM {
	if in == nil {
		return nil
	}
	out := new(RoleAssignment_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignments_Spec) DeepCopyInto(out *RoleAssignments_Spec) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.ConditionVersion != nil {
		in, out := &in.ConditionVersion, &out.ConditionVersion
		*out = new(string)
		**out = **in
	}
	if in.DelegatedManagedIdentityResourceId != nil {
		in, out := &in.DelegatedManagedIdentityResourceId, &out.DelegatedManagedIdentityResourceId
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	out.Owner = in.Owner
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(RoleAssignmentPropertiesPrincipalType)
		**out = **in
	}
	out.RoleDefinitionReference = in.RoleDefinitionReference
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignments_Spec.
func (in *RoleAssignments_Spec) DeepCopy() *RoleAssignments_Spec {
	if in == nil {
		return nil
	}
	out := new(RoleAssignments_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignments_SpecARM) DeepCopyInto(out *RoleAssignments_SpecARM) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	in.Properties.DeepCopyInto(&out.Properties)
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignments_SpecARM.
func (in *RoleAssignments_SpecARM) DeepCopy() *RoleAssignments_SpecARM {
	if in == nil {
		return nil
	}
	out := new(RoleAssignments_SpecARM)
	in.DeepCopyInto(out)
	return out
}
