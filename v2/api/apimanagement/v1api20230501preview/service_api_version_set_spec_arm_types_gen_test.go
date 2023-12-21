// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

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

func Test_Service_ApiVersionSet_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_ApiVersionSet_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_ApiVersionSet_Spec_ARM, Service_ApiVersionSet_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_ApiVersionSet_Spec_ARM runs a test to see if a specific instance of Service_ApiVersionSet_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_ApiVersionSet_Spec_ARM(subject Service_ApiVersionSet_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_ApiVersionSet_Spec_ARM
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

// Generator of Service_ApiVersionSet_Spec_ARM instances for property testing - lazily instantiated by
// Service_ApiVersionSet_Spec_ARMGenerator()
var service_ApiVersionSet_Spec_ARMGenerator gopter.Gen

// Service_ApiVersionSet_Spec_ARMGenerator returns a generator of Service_ApiVersionSet_Spec_ARM instances for property testing.
// We first initialize service_ApiVersionSet_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_ApiVersionSet_Spec_ARMGenerator() gopter.Gen {
	if service_ApiVersionSet_Spec_ARMGenerator != nil {
		return service_ApiVersionSet_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_ApiVersionSet_Spec_ARM(generators)
	service_ApiVersionSet_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_ApiVersionSet_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_ApiVersionSet_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForService_ApiVersionSet_Spec_ARM(generators)
	service_ApiVersionSet_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_ApiVersionSet_Spec_ARM{}), generators)

	return service_ApiVersionSet_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_ApiVersionSet_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_ApiVersionSet_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForService_ApiVersionSet_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_ApiVersionSet_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ApiVersionSetContractProperties_ARMGenerator())
}

func Test_ApiVersionSetContractProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiVersionSetContractProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiVersionSetContractProperties_ARM, ApiVersionSetContractProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiVersionSetContractProperties_ARM runs a test to see if a specific instance of ApiVersionSetContractProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApiVersionSetContractProperties_ARM(subject ApiVersionSetContractProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiVersionSetContractProperties_ARM
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

// Generator of ApiVersionSetContractProperties_ARM instances for property testing - lazily instantiated by
// ApiVersionSetContractProperties_ARMGenerator()
var apiVersionSetContractProperties_ARMGenerator gopter.Gen

// ApiVersionSetContractProperties_ARMGenerator returns a generator of ApiVersionSetContractProperties_ARM instances for property testing.
func ApiVersionSetContractProperties_ARMGenerator() gopter.Gen {
	if apiVersionSetContractProperties_ARMGenerator != nil {
		return apiVersionSetContractProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiVersionSetContractProperties_ARM(generators)
	apiVersionSetContractProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ApiVersionSetContractProperties_ARM{}), generators)

	return apiVersionSetContractProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApiVersionSetContractProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiVersionSetContractProperties_ARM(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["VersionHeaderName"] = gen.PtrOf(gen.AlphaString())
	gens["VersionQueryName"] = gen.PtrOf(gen.AlphaString())
	gens["VersioningScheme"] = gen.PtrOf(gen.OneConstOf(ApiVersionSetContractProperties_VersioningScheme_Header, ApiVersionSetContractProperties_VersioningScheme_Query, ApiVersionSetContractProperties_VersioningScheme_Segment))
}
