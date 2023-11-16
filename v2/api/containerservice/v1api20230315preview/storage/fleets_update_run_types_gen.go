// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=containerservice.azure.com,resources=fleetsupdateruns,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=containerservice.azure.com,resources={fleetsupdateruns/status,fleetsupdateruns/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230315preview.FleetsUpdateRun
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateRuns/{updateRunName}
type FleetsUpdateRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Fleets_UpdateRun_Spec   `json:"spec,omitempty"`
	Status            Fleets_UpdateRun_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FleetsUpdateRun{}

// GetConditions returns the conditions of the resource
func (updateRun *FleetsUpdateRun) GetConditions() conditions.Conditions {
	return updateRun.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (updateRun *FleetsUpdateRun) SetConditions(conditions conditions.Conditions) {
	updateRun.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &FleetsUpdateRun{}

// AzureName returns the Azure name of the resource
func (updateRun *FleetsUpdateRun) AzureName() string {
	return updateRun.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-03-15-preview"
func (updateRun FleetsUpdateRun) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (updateRun *FleetsUpdateRun) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (updateRun *FleetsUpdateRun) GetSpec() genruntime.ConvertibleSpec {
	return &updateRun.Spec
}

// GetStatus returns the status of this resource
func (updateRun *FleetsUpdateRun) GetStatus() genruntime.ConvertibleStatus {
	return &updateRun.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (updateRun *FleetsUpdateRun) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/fleets/updateRuns"
func (updateRun *FleetsUpdateRun) GetType() string {
	return "Microsoft.ContainerService/fleets/updateRuns"
}

// NewEmptyStatus returns a new empty (blank) status
func (updateRun *FleetsUpdateRun) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Fleets_UpdateRun_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (updateRun *FleetsUpdateRun) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(updateRun.Spec)
	return updateRun.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (updateRun *FleetsUpdateRun) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Fleets_UpdateRun_STATUS); ok {
		updateRun.Status = *st
		return nil
	}

	// Convert status to required version
	var st Fleets_UpdateRun_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	updateRun.Status = st
	return nil
}

// Hub marks that this FleetsUpdateRun is the hub type for conversion
func (updateRun *FleetsUpdateRun) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (updateRun *FleetsUpdateRun) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: updateRun.Spec.OriginalVersion,
		Kind:    "FleetsUpdateRun",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230315preview.FleetsUpdateRun
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateRuns/{updateRunName}
type FleetsUpdateRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FleetsUpdateRun `json:"items"`
}

// Storage version of v1api20230315preview.Fleets_UpdateRun_Spec
type Fleets_UpdateRun_Spec struct {
	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern="^[a-z0-9]([-a-z0-9]*[a-z0-9])?$"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName            string                `json:"azureName,omitempty"`
	ManagedClusterUpdate *ManagedClusterUpdate `json:"managedClusterUpdate,omitempty"`
	OriginalVersion      string                `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a containerservice.azure.com/Fleet resource
	Owner       *genruntime.KnownResourceReference `group:"containerservice.azure.com" json:"owner,omitempty" kind:"Fleet"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Strategy    *UpdateRunStrategy                 `json:"strategy,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Fleets_UpdateRun_Spec{}

// ConvertSpecFrom populates our Fleets_UpdateRun_Spec from the provided source
func (updateRun *Fleets_UpdateRun_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == updateRun {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(updateRun)
}

// ConvertSpecTo populates the provided destination from our Fleets_UpdateRun_Spec
func (updateRun *Fleets_UpdateRun_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == updateRun {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(updateRun)
}

// Storage version of v1api20230315preview.Fleets_UpdateRun_STATUS
type Fleets_UpdateRun_STATUS struct {
	Conditions           []conditions.Condition       `json:"conditions,omitempty"`
	ETag                 *string                      `json:"eTag,omitempty"`
	Id                   *string                      `json:"id,omitempty"`
	ManagedClusterUpdate *ManagedClusterUpdate_STATUS `json:"managedClusterUpdate,omitempty"`
	Name                 *string                      `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	ProvisioningState    *string                      `json:"provisioningState,omitempty"`
	Status               *UpdateRunStatus_STATUS      `json:"status,omitempty"`
	Strategy             *UpdateRunStrategy_STATUS    `json:"strategy,omitempty"`
	SystemData           *SystemData_STATUS           `json:"systemData,omitempty"`
	Type                 *string                      `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Fleets_UpdateRun_STATUS{}

// ConvertStatusFrom populates our Fleets_UpdateRun_STATUS from the provided source
func (updateRun *Fleets_UpdateRun_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == updateRun {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(updateRun)
}

// ConvertStatusTo populates the provided destination from our Fleets_UpdateRun_STATUS
func (updateRun *Fleets_UpdateRun_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == updateRun {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(updateRun)
}

// Storage version of v1api20230315preview.ManagedClusterUpdate
// The update to be applied to the ManagedClusters.
type ManagedClusterUpdate struct {
	PropertyBag genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Upgrade     *ManagedClusterUpgradeSpec `json:"upgrade,omitempty"`
}

// Storage version of v1api20230315preview.ManagedClusterUpdate_STATUS
// The update to be applied to the ManagedClusters.
type ManagedClusterUpdate_STATUS struct {
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Upgrade     *ManagedClusterUpgradeSpec_STATUS `json:"upgrade,omitempty"`
}

// Storage version of v1api20230315preview.UpdateRunStatus_STATUS
// The status of a UpdateRun.
type UpdateRunStatus_STATUS struct {
	PropertyBag genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Stages      []UpdateStageStatus_STATUS `json:"stages,omitempty"`
	Status      *UpdateStatus_STATUS       `json:"status,omitempty"`
}

// Storage version of v1api20230315preview.UpdateRunStrategy
// Defines the update sequence of the clusters via stages and groups.
// Stages within a run are executed sequentially one
// after another.
// Groups within a stage are executed in parallel.
// Member clusters within a group are updated sequentially
// one after another.
// A valid strategy contains no duplicate groups within or across stages.
type UpdateRunStrategy struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Stages      []UpdateStage          `json:"stages,omitempty"`
}

// Storage version of v1api20230315preview.UpdateRunStrategy_STATUS
// Defines the update sequence of the clusters via stages and groups.
// Stages within a run are executed sequentially one
// after another.
// Groups within a stage are executed in parallel.
// Member clusters within a group are updated sequentially
// one after another.
// A valid strategy contains no duplicate groups within or across stages.
type UpdateRunStrategy_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Stages      []UpdateStage_STATUS   `json:"stages,omitempty"`
}

// Storage version of v1api20230315preview.ManagedClusterUpgradeSpec
// The upgrade to apply to a ManagedCluster.
type ManagedClusterUpgradeSpec struct {
	KubernetesVersion *string                `json:"kubernetesVersion,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

// Storage version of v1api20230315preview.ManagedClusterUpgradeSpec_STATUS
// The upgrade to apply to a ManagedCluster.
type ManagedClusterUpgradeSpec_STATUS struct {
	KubernetesVersion *string                `json:"kubernetesVersion,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

// Storage version of v1api20230315preview.UpdateStage
// Defines a stage which contains the groups to update and the steps to take (e.g., wait for a time period) before starting
// the next stage.
type UpdateStage struct {
	AfterStageWaitInSeconds *int                   `json:"afterStageWaitInSeconds,omitempty"`
	Groups                  []UpdateGroup          `json:"groups,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230315preview.UpdateStage_STATUS
// Defines a stage which contains the groups to update and the steps to take (e.g., wait for a time period) before starting
// the next stage.
type UpdateStage_STATUS struct {
	AfterStageWaitInSeconds *int                   `json:"afterStageWaitInSeconds,omitempty"`
	Groups                  []UpdateGroup_STATUS   `json:"groups,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230315preview.UpdateStageStatus_STATUS
// The status of a UpdateStage.
type UpdateStageStatus_STATUS struct {
	AfterStageWaitStatus *WaitStatus_STATUS         `json:"afterStageWaitStatus,omitempty"`
	Groups               []UpdateGroupStatus_STATUS `json:"groups,omitempty"`
	Name                 *string                    `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Status               *UpdateStatus_STATUS       `json:"status,omitempty"`
}

// Storage version of v1api20230315preview.UpdateStatus_STATUS
// The status for an operation or group of operations.
type UpdateStatus_STATUS struct {
	CompletedTime *string                `json:"completedTime,omitempty"`
	Error         *ErrorDetail_STATUS    `json:"error,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartTime     *string                `json:"startTime,omitempty"`
	State         *string                `json:"state,omitempty"`
}

// Storage version of v1api20230315preview.ErrorDetail_STATUS
// The error detail.
type ErrorDetail_STATUS struct {
	AdditionalInfo []ErrorAdditionalInfo_STATUS  `json:"additionalInfo,omitempty"`
	Code           *string                       `json:"code,omitempty"`
	Details        []ErrorDetail_STATUS_Unrolled `json:"details,omitempty"`
	Message        *string                       `json:"message,omitempty"`
	PropertyBag    genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Target         *string                       `json:"target,omitempty"`
}

// Storage version of v1api20230315preview.UpdateGroup
// A group to be updated.
type UpdateGroup struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230315preview.UpdateGroup_STATUS
// A group to be updated.
type UpdateGroup_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230315preview.UpdateGroupStatus_STATUS
// The status of a UpdateGroup.
type UpdateGroupStatus_STATUS struct {
	Members     []MemberUpdateStatus_STATUS `json:"members,omitempty"`
	Name        *string                     `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	Status      *UpdateStatus_STATUS        `json:"status,omitempty"`
}

// Storage version of v1api20230315preview.WaitStatus_STATUS
// The status of the wait duration.
type WaitStatus_STATUS struct {
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status                *UpdateStatus_STATUS   `json:"status,omitempty"`
	WaitDurationInSeconds *int                   `json:"waitDurationInSeconds,omitempty"`
}

// Storage version of v1api20230315preview.ErrorAdditionalInfo_STATUS
// The resource management error additional info.
type ErrorAdditionalInfo_STATUS struct {
	Info        map[string]v1.JSON     `json:"info,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20230315preview.ErrorDetail_STATUS_Unrolled
type ErrorDetail_STATUS_Unrolled struct {
	AdditionalInfo []ErrorAdditionalInfo_STATUS `json:"additionalInfo,omitempty"`
	Code           *string                      `json:"code,omitempty"`
	Message        *string                      `json:"message,omitempty"`
	PropertyBag    genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Target         *string                      `json:"target,omitempty"`
}

// Storage version of v1api20230315preview.MemberUpdateStatus_STATUS
// The status of a member update operation.
type MemberUpdateStatus_STATUS struct {
	ClusterResourceId *string                `json:"clusterResourceId,omitempty"`
	Name              *string                `json:"name,omitempty"`
	OperationId       *string                `json:"operationId,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status            *UpdateStatus_STATUS   `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&FleetsUpdateRun{}, &FleetsUpdateRunList{})
}