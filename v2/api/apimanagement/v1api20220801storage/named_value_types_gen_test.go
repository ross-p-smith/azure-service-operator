// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801storage

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

func Test_NamedValue_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamedValue via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamedValue, NamedValueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamedValue runs a test to see if a specific instance of NamedValue round trips to JSON and back losslessly
func RunJSONSerializationTestForNamedValue(subject NamedValue) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamedValue
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

// Generator of NamedValue instances for property testing - lazily instantiated by NamedValueGenerator()
var namedValueGenerator gopter.Gen

// NamedValueGenerator returns a generator of NamedValue instances for property testing.
func NamedValueGenerator() gopter.Gen {
	if namedValueGenerator != nil {
		return namedValueGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamedValue(generators)
	namedValueGenerator = gen.Struct(reflect.TypeOf(NamedValue{}), generators)

	return namedValueGenerator
}

// AddRelatedPropertyGeneratorsForNamedValue is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamedValue(gens map[string]gopter.Gen) {
	gens["Spec"] = Service_NamedValue_SpecGenerator()
	gens["Status"] = Service_NamedValue_STATUSGenerator()
}

func Test_Service_NamedValue_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_NamedValue_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_NamedValue_Spec, Service_NamedValue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_NamedValue_Spec runs a test to see if a specific instance of Service_NamedValue_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForService_NamedValue_Spec(subject Service_NamedValue_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_NamedValue_Spec
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

// Generator of Service_NamedValue_Spec instances for property testing - lazily instantiated by
// Service_NamedValue_SpecGenerator()
var service_NamedValue_SpecGenerator gopter.Gen

// Service_NamedValue_SpecGenerator returns a generator of Service_NamedValue_Spec instances for property testing.
// We first initialize service_NamedValue_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_NamedValue_SpecGenerator() gopter.Gen {
	if service_NamedValue_SpecGenerator != nil {
		return service_NamedValue_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_NamedValue_Spec(generators)
	service_NamedValue_SpecGenerator = gen.Struct(reflect.TypeOf(Service_NamedValue_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_NamedValue_Spec(generators)
	AddRelatedPropertyGeneratorsForService_NamedValue_Spec(generators)
	service_NamedValue_SpecGenerator = gen.Struct(reflect.TypeOf(Service_NamedValue_Spec{}), generators)

	return service_NamedValue_SpecGenerator
}

// AddIndependentPropertyGeneratorsForService_NamedValue_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_NamedValue_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Secret"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.SliceOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForService_NamedValue_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_NamedValue_Spec(gens map[string]gopter.Gen) {
	gens["KeyVault"] = gen.PtrOf(KeyVaultContractCreatePropertiesGenerator())
}

func Test_Service_NamedValue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_NamedValue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_NamedValue_STATUS, Service_NamedValue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_NamedValue_STATUS runs a test to see if a specific instance of Service_NamedValue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForService_NamedValue_STATUS(subject Service_NamedValue_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_NamedValue_STATUS
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

// Generator of Service_NamedValue_STATUS instances for property testing - lazily instantiated by
// Service_NamedValue_STATUSGenerator()
var service_NamedValue_STATUSGenerator gopter.Gen

// Service_NamedValue_STATUSGenerator returns a generator of Service_NamedValue_STATUS instances for property testing.
// We first initialize service_NamedValue_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_NamedValue_STATUSGenerator() gopter.Gen {
	if service_NamedValue_STATUSGenerator != nil {
		return service_NamedValue_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_NamedValue_STATUS(generators)
	service_NamedValue_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_NamedValue_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_NamedValue_STATUS(generators)
	AddRelatedPropertyGeneratorsForService_NamedValue_STATUS(generators)
	service_NamedValue_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_NamedValue_STATUS{}), generators)

	return service_NamedValue_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForService_NamedValue_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_NamedValue_STATUS(gens map[string]gopter.Gen) {
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Secret"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForService_NamedValue_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_NamedValue_STATUS(gens map[string]gopter.Gen) {
	gens["KeyVault"] = gen.PtrOf(KeyVaultContractProperties_STATUSGenerator())
}

func Test_KeyVaultContractCreateProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultContractCreateProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultContractCreateProperties, KeyVaultContractCreatePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultContractCreateProperties runs a test to see if a specific instance of KeyVaultContractCreateProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultContractCreateProperties(subject KeyVaultContractCreateProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultContractCreateProperties
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

// Generator of KeyVaultContractCreateProperties instances for property testing - lazily instantiated by
// KeyVaultContractCreatePropertiesGenerator()
var keyVaultContractCreatePropertiesGenerator gopter.Gen

// KeyVaultContractCreatePropertiesGenerator returns a generator of KeyVaultContractCreateProperties instances for property testing.
func KeyVaultContractCreatePropertiesGenerator() gopter.Gen {
	if keyVaultContractCreatePropertiesGenerator != nil {
		return keyVaultContractCreatePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultContractCreateProperties(generators)
	keyVaultContractCreatePropertiesGenerator = gen.Struct(reflect.TypeOf(KeyVaultContractCreateProperties{}), generators)

	return keyVaultContractCreatePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultContractCreateProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultContractCreateProperties(gens map[string]gopter.Gen) {
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultContractProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultContractProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultContractProperties_STATUS, KeyVaultContractProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultContractProperties_STATUS runs a test to see if a specific instance of KeyVaultContractProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultContractProperties_STATUS(subject KeyVaultContractProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultContractProperties_STATUS
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

// Generator of KeyVaultContractProperties_STATUS instances for property testing - lazily instantiated by
// KeyVaultContractProperties_STATUSGenerator()
var keyVaultContractProperties_STATUSGenerator gopter.Gen

// KeyVaultContractProperties_STATUSGenerator returns a generator of KeyVaultContractProperties_STATUS instances for property testing.
// We first initialize keyVaultContractProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultContractProperties_STATUSGenerator() gopter.Gen {
	if keyVaultContractProperties_STATUSGenerator != nil {
		return keyVaultContractProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultContractProperties_STATUS(generators)
	keyVaultContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyVaultContractProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultContractProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForKeyVaultContractProperties_STATUS(generators)
	keyVaultContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyVaultContractProperties_STATUS{}), generators)

	return keyVaultContractProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultContractProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
	gens["SecretIdentifier"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultContractProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["LastStatus"] = gen.PtrOf(KeyVaultLastAccessStatusContractProperties_STATUSGenerator())
}

func Test_KeyVaultLastAccessStatusContractProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultLastAccessStatusContractProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultLastAccessStatusContractProperties_STATUS, KeyVaultLastAccessStatusContractProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultLastAccessStatusContractProperties_STATUS runs a test to see if a specific instance of KeyVaultLastAccessStatusContractProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultLastAccessStatusContractProperties_STATUS(subject KeyVaultLastAccessStatusContractProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultLastAccessStatusContractProperties_STATUS
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

// Generator of KeyVaultLastAccessStatusContractProperties_STATUS instances for property testing - lazily instantiated
// by KeyVaultLastAccessStatusContractProperties_STATUSGenerator()
var keyVaultLastAccessStatusContractProperties_STATUSGenerator gopter.Gen

// KeyVaultLastAccessStatusContractProperties_STATUSGenerator returns a generator of KeyVaultLastAccessStatusContractProperties_STATUS instances for property testing.
func KeyVaultLastAccessStatusContractProperties_STATUSGenerator() gopter.Gen {
	if keyVaultLastAccessStatusContractProperties_STATUSGenerator != nil {
		return keyVaultLastAccessStatusContractProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultLastAccessStatusContractProperties_STATUS(generators)
	keyVaultLastAccessStatusContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyVaultLastAccessStatusContractProperties_STATUS{}), generators)

	return keyVaultLastAccessStatusContractProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultLastAccessStatusContractProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultLastAccessStatusContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["TimeStampUtc"] = gen.PtrOf(gen.AlphaString())
}
