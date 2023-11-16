// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Extension_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Extension_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtension_STATUS_ARM, Extension_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtension_STATUS_ARM runs a test to see if a specific instance of Extension_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtension_STATUS_ARM(subject Extension_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Extension_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Extension_STATUS_ARM instances for property testing - lazily instantiated by
// Extension_STATUS_ARMGenerator()
var extension_STATUS_ARMGenerator gopter.Gen

// Extension_STATUS_ARMGenerator returns a generator of Extension_STATUS_ARM instances for property testing.
// We first initialize extension_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Extension_STATUS_ARMGenerator() gopter.Gen {
	if extension_STATUS_ARMGenerator != nil {
		return extension_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtension_STATUS_ARM(generators)
	extension_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Extension_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtension_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForExtension_STATUS_ARM(generators)
	extension_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Extension_STATUS_ARM{}), generators)

	return extension_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExtension_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtension_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForExtension_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForExtension_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUS_ARMGenerator())
	gens["Plan"] = gen.PtrOf(Plan_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(Extension_Properties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_Extension_Properties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Extension_Properties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtension_Properties_STATUS_ARM, Extension_Properties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtension_Properties_STATUS_ARM runs a test to see if a specific instance of Extension_Properties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtension_Properties_STATUS_ARM(subject Extension_Properties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Extension_Properties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Extension_Properties_STATUS_ARM instances for property testing - lazily instantiated by
// Extension_Properties_STATUS_ARMGenerator()
var extension_Properties_STATUS_ARMGenerator gopter.Gen

// Extension_Properties_STATUS_ARMGenerator returns a generator of Extension_Properties_STATUS_ARM instances for property testing.
// We first initialize extension_Properties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Extension_Properties_STATUS_ARMGenerator() gopter.Gen {
	if extension_Properties_STATUS_ARMGenerator != nil {
		return extension_Properties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtension_Properties_STATUS_ARM(generators)
	extension_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Extension_Properties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtension_Properties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForExtension_Properties_STATUS_ARM(generators)
	extension_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Extension_Properties_STATUS_ARM{}), generators)

	return extension_Properties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExtension_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtension_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["ConfigurationSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["CurrentVersion"] = gen.PtrOf(gen.AlphaString())
	gens["CustomLocationSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["ExtensionType"] = gen.PtrOf(gen.AlphaString())
	gens["IsSystemExtension"] = gen.PtrOf(gen.Bool())
	gens["PackageUri"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningStateDefinition_STATUS_Canceled,
		ProvisioningStateDefinition_STATUS_Creating,
		ProvisioningStateDefinition_STATUS_Deleting,
		ProvisioningStateDefinition_STATUS_Failed,
		ProvisioningStateDefinition_STATUS_Succeeded,
		ProvisioningStateDefinition_STATUS_Updating))
	gens["ReleaseTrain"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForExtension_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForExtension_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AksAssignedIdentity"] = gen.PtrOf(Extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator())
	gens["ErrorInfo"] = gen.PtrOf(ErrorDetail_STATUS_ARMGenerator())
	gens["Scope"] = gen.PtrOf(Scope_STATUS_ARMGenerator())
	gens["Statuses"] = gen.SliceOf(ExtensionStatus_STATUS_ARMGenerator())
}

func Test_Identity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_STATUS_ARM, Identity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_STATUS_ARM runs a test to see if a specific instance of Identity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_STATUS_ARM(subject Identity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Identity_STATUS_ARM instances for property testing - lazily instantiated by
// Identity_STATUS_ARMGenerator()
var identity_STATUS_ARMGenerator gopter.Gen

// Identity_STATUS_ARMGenerator returns a generator of Identity_STATUS_ARM instances for property testing.
func Identity_STATUS_ARMGenerator() gopter.Gen {
	if identity_STATUS_ARMGenerator != nil {
		return identity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(generators)
	identity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUS_ARM{}), generators)

	return identity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(Identity_Type_STATUS_SystemAssigned))
}

func Test_Plan_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Plan_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPlan_STATUS_ARM, Plan_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPlan_STATUS_ARM runs a test to see if a specific instance of Plan_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPlan_STATUS_ARM(subject Plan_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Plan_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Plan_STATUS_ARM instances for property testing - lazily instantiated by Plan_STATUS_ARMGenerator()
var plan_STATUS_ARMGenerator gopter.Gen

// Plan_STATUS_ARMGenerator returns a generator of Plan_STATUS_ARM instances for property testing.
func Plan_STATUS_ARMGenerator() gopter.Gen {
	if plan_STATUS_ARMGenerator != nil {
		return plan_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPlan_STATUS_ARM(generators)
	plan_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Plan_STATUS_ARM{}), generators)

	return plan_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPlan_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPlan_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Product"] = gen.PtrOf(gen.AlphaString())
	gens["PromotionCode"] = gen.PtrOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS_ARM, SystemData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS_ARM runs a test to see if a specific instance of SystemData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS_ARM(subject SystemData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SystemData_STATUS_ARM instances for property testing - lazily instantiated by
// SystemData_STATUS_ARMGenerator()
var systemData_STATUS_ARMGenerator gopter.Gen

// SystemData_STATUS_ARMGenerator returns a generator of SystemData_STATUS_ARM instances for property testing.
func SystemData_STATUS_ARMGenerator() gopter.Gen {
	if systemData_STATUS_ARMGenerator != nil {
		return systemData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(generators)
	systemData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS_ARM{}), generators)

	return systemData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}

func Test_ErrorDetail_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorDetail_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorDetail_STATUS_ARM, ErrorDetail_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorDetail_STATUS_ARM runs a test to see if a specific instance of ErrorDetail_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorDetail_STATUS_ARM(subject ErrorDetail_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorDetail_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ErrorDetail_STATUS_ARM instances for property testing - lazily instantiated by
// ErrorDetail_STATUS_ARMGenerator()
var errorDetail_STATUS_ARMGenerator gopter.Gen

// ErrorDetail_STATUS_ARMGenerator returns a generator of ErrorDetail_STATUS_ARM instances for property testing.
// We first initialize errorDetail_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ErrorDetail_STATUS_ARMGenerator() gopter.Gen {
	if errorDetail_STATUS_ARMGenerator != nil {
		return errorDetail_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetail_STATUS_ARM(generators)
	errorDetail_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetail_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForErrorDetail_STATUS_ARM(generators)
	errorDetail_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_STATUS_ARM{}), generators)

	return errorDetail_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForErrorDetail_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForErrorDetail_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForErrorDetail_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForErrorDetail_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AdditionalInfo"] = gen.SliceOf(ErrorAdditionalInfo_STATUS_ARMGenerator())
	gens["Details"] = gen.SliceOf(ErrorDetail_STATUS_Unrolled_ARMGenerator())
}

func Test_Extension_Properties_AksAssignedIdentity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Extension_Properties_AksAssignedIdentity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtension_Properties_AksAssignedIdentity_STATUS_ARM, Extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtension_Properties_AksAssignedIdentity_STATUS_ARM runs a test to see if a specific instance of Extension_Properties_AksAssignedIdentity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtension_Properties_AksAssignedIdentity_STATUS_ARM(subject Extension_Properties_AksAssignedIdentity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Extension_Properties_AksAssignedIdentity_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Extension_Properties_AksAssignedIdentity_STATUS_ARM instances for property testing - lazily instantiated
// by Extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator()
var extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator gopter.Gen

// Extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator returns a generator of Extension_Properties_AksAssignedIdentity_STATUS_ARM instances for property testing.
func Extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator() gopter.Gen {
	if extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator != nil {
		return extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtension_Properties_AksAssignedIdentity_STATUS_ARM(generators)
	extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Extension_Properties_AksAssignedIdentity_STATUS_ARM{}), generators)

	return extension_Properties_AksAssignedIdentity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExtension_Properties_AksAssignedIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtension_Properties_AksAssignedIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(Extension_Properties_AksAssignedIdentity_Type_STATUS_SystemAssigned, Extension_Properties_AksAssignedIdentity_Type_STATUS_UserAssigned))
}

func Test_ExtensionStatus_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtensionStatus_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtensionStatus_STATUS_ARM, ExtensionStatus_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtensionStatus_STATUS_ARM runs a test to see if a specific instance of ExtensionStatus_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtensionStatus_STATUS_ARM(subject ExtensionStatus_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtensionStatus_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ExtensionStatus_STATUS_ARM instances for property testing - lazily instantiated by
// ExtensionStatus_STATUS_ARMGenerator()
var extensionStatus_STATUS_ARMGenerator gopter.Gen

// ExtensionStatus_STATUS_ARMGenerator returns a generator of ExtensionStatus_STATUS_ARM instances for property testing.
func ExtensionStatus_STATUS_ARMGenerator() gopter.Gen {
	if extensionStatus_STATUS_ARMGenerator != nil {
		return extensionStatus_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtensionStatus_STATUS_ARM(generators)
	extensionStatus_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ExtensionStatus_STATUS_ARM{}), generators)

	return extensionStatus_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExtensionStatus_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtensionStatus_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Level"] = gen.PtrOf(gen.OneConstOf(ExtensionStatus_Level_STATUS_Error, ExtensionStatus_Level_STATUS_Information, ExtensionStatus_Level_STATUS_Warning))
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Time"] = gen.PtrOf(gen.AlphaString())
}

func Test_Scope_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Scope_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScope_STATUS_ARM, Scope_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScope_STATUS_ARM runs a test to see if a specific instance of Scope_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScope_STATUS_ARM(subject Scope_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Scope_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Scope_STATUS_ARM instances for property testing - lazily instantiated by Scope_STATUS_ARMGenerator()
var scope_STATUS_ARMGenerator gopter.Gen

// Scope_STATUS_ARMGenerator returns a generator of Scope_STATUS_ARM instances for property testing.
func Scope_STATUS_ARMGenerator() gopter.Gen {
	if scope_STATUS_ARMGenerator != nil {
		return scope_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScope_STATUS_ARM(generators)
	scope_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Scope_STATUS_ARM{}), generators)

	return scope_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForScope_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScope_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Cluster"] = gen.PtrOf(ScopeCluster_STATUS_ARMGenerator())
	gens["Namespace"] = gen.PtrOf(ScopeNamespace_STATUS_ARMGenerator())
}

func Test_ErrorAdditionalInfo_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorAdditionalInfo_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorAdditionalInfo_STATUS_ARM, ErrorAdditionalInfo_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorAdditionalInfo_STATUS_ARM runs a test to see if a specific instance of ErrorAdditionalInfo_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorAdditionalInfo_STATUS_ARM(subject ErrorAdditionalInfo_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorAdditionalInfo_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ErrorAdditionalInfo_STATUS_ARM instances for property testing - lazily instantiated by
// ErrorAdditionalInfo_STATUS_ARMGenerator()
var errorAdditionalInfo_STATUS_ARMGenerator gopter.Gen

// ErrorAdditionalInfo_STATUS_ARMGenerator returns a generator of ErrorAdditionalInfo_STATUS_ARM instances for property testing.
func ErrorAdditionalInfo_STATUS_ARMGenerator() gopter.Gen {
	if errorAdditionalInfo_STATUS_ARMGenerator != nil {
		return errorAdditionalInfo_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorAdditionalInfo_STATUS_ARM(generators)
	errorAdditionalInfo_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ErrorAdditionalInfo_STATUS_ARM{}), generators)

	return errorAdditionalInfo_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForErrorAdditionalInfo_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForErrorAdditionalInfo_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ErrorDetail_STATUS_Unrolled_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorDetail_STATUS_Unrolled_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorDetail_STATUS_Unrolled_ARM, ErrorDetail_STATUS_Unrolled_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorDetail_STATUS_Unrolled_ARM runs a test to see if a specific instance of ErrorDetail_STATUS_Unrolled_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorDetail_STATUS_Unrolled_ARM(subject ErrorDetail_STATUS_Unrolled_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorDetail_STATUS_Unrolled_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ErrorDetail_STATUS_Unrolled_ARM instances for property testing - lazily instantiated by
// ErrorDetail_STATUS_Unrolled_ARMGenerator()
var errorDetail_STATUS_Unrolled_ARMGenerator gopter.Gen

// ErrorDetail_STATUS_Unrolled_ARMGenerator returns a generator of ErrorDetail_STATUS_Unrolled_ARM instances for property testing.
// We first initialize errorDetail_STATUS_Unrolled_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ErrorDetail_STATUS_Unrolled_ARMGenerator() gopter.Gen {
	if errorDetail_STATUS_Unrolled_ARMGenerator != nil {
		return errorDetail_STATUS_Unrolled_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetail_STATUS_Unrolled_ARM(generators)
	errorDetail_STATUS_Unrolled_ARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_STATUS_Unrolled_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetail_STATUS_Unrolled_ARM(generators)
	AddRelatedPropertyGeneratorsForErrorDetail_STATUS_Unrolled_ARM(generators)
	errorDetail_STATUS_Unrolled_ARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_STATUS_Unrolled_ARM{}), generators)

	return errorDetail_STATUS_Unrolled_ARMGenerator
}

// AddIndependentPropertyGeneratorsForErrorDetail_STATUS_Unrolled_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForErrorDetail_STATUS_Unrolled_ARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForErrorDetail_STATUS_Unrolled_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForErrorDetail_STATUS_Unrolled_ARM(gens map[string]gopter.Gen) {
	gens["AdditionalInfo"] = gen.SliceOf(ErrorAdditionalInfo_STATUS_ARMGenerator())
}

func Test_ScopeCluster_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScopeCluster_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScopeCluster_STATUS_ARM, ScopeCluster_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScopeCluster_STATUS_ARM runs a test to see if a specific instance of ScopeCluster_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScopeCluster_STATUS_ARM(subject ScopeCluster_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScopeCluster_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ScopeCluster_STATUS_ARM instances for property testing - lazily instantiated by
// ScopeCluster_STATUS_ARMGenerator()
var scopeCluster_STATUS_ARMGenerator gopter.Gen

// ScopeCluster_STATUS_ARMGenerator returns a generator of ScopeCluster_STATUS_ARM instances for property testing.
func ScopeCluster_STATUS_ARMGenerator() gopter.Gen {
	if scopeCluster_STATUS_ARMGenerator != nil {
		return scopeCluster_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScopeCluster_STATUS_ARM(generators)
	scopeCluster_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ScopeCluster_STATUS_ARM{}), generators)

	return scopeCluster_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForScopeCluster_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScopeCluster_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ReleaseNamespace"] = gen.PtrOf(gen.AlphaString())
}

func Test_ScopeNamespace_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScopeNamespace_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScopeNamespace_STATUS_ARM, ScopeNamespace_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScopeNamespace_STATUS_ARM runs a test to see if a specific instance of ScopeNamespace_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScopeNamespace_STATUS_ARM(subject ScopeNamespace_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScopeNamespace_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ScopeNamespace_STATUS_ARM instances for property testing - lazily instantiated by
// ScopeNamespace_STATUS_ARMGenerator()
var scopeNamespace_STATUS_ARMGenerator gopter.Gen

// ScopeNamespace_STATUS_ARMGenerator returns a generator of ScopeNamespace_STATUS_ARM instances for property testing.
func ScopeNamespace_STATUS_ARMGenerator() gopter.Gen {
	if scopeNamespace_STATUS_ARMGenerator != nil {
		return scopeNamespace_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScopeNamespace_STATUS_ARM(generators)
	scopeNamespace_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ScopeNamespace_STATUS_ARM{}), generators)

	return scopeNamespace_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForScopeNamespace_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScopeNamespace_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["TargetNamespace"] = gen.PtrOf(gen.AlphaString())
}