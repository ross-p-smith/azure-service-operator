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

func Test_Service_Policy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Policy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Policy_STATUS_ARM, Service_Policy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Policy_STATUS_ARM runs a test to see if a specific instance of Service_Policy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Policy_STATUS_ARM(subject Service_Policy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Policy_STATUS_ARM
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

// Generator of Service_Policy_STATUS_ARM instances for property testing - lazily instantiated by
// Service_Policy_STATUS_ARMGenerator()
var service_Policy_STATUS_ARMGenerator gopter.Gen

// Service_Policy_STATUS_ARMGenerator returns a generator of Service_Policy_STATUS_ARM instances for property testing.
// We first initialize service_Policy_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_Policy_STATUS_ARMGenerator() gopter.Gen {
	if service_Policy_STATUS_ARMGenerator != nil {
		return service_Policy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Policy_STATUS_ARM(generators)
	service_Policy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_Policy_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Policy_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForService_Policy_STATUS_ARM(generators)
	service_Policy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_Policy_STATUS_ARM{}), generators)

	return service_Policy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_Policy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Policy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForService_Policy_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_Policy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PolicyContractProperties_STATUS_ARMGenerator())
}

func Test_PolicyContractProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PolicyContractProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicyContractProperties_STATUS_ARM, PolicyContractProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicyContractProperties_STATUS_ARM runs a test to see if a specific instance of PolicyContractProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicyContractProperties_STATUS_ARM(subject PolicyContractProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PolicyContractProperties_STATUS_ARM
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

// Generator of PolicyContractProperties_STATUS_ARM instances for property testing - lazily instantiated by
// PolicyContractProperties_STATUS_ARMGenerator()
var policyContractProperties_STATUS_ARMGenerator gopter.Gen

// PolicyContractProperties_STATUS_ARMGenerator returns a generator of PolicyContractProperties_STATUS_ARM instances for property testing.
func PolicyContractProperties_STATUS_ARMGenerator() gopter.Gen {
	if policyContractProperties_STATUS_ARMGenerator != nil {
		return policyContractProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPolicyContractProperties_STATUS_ARM(generators)
	policyContractProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PolicyContractProperties_STATUS_ARM{}), generators)

	return policyContractProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPolicyContractProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPolicyContractProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Format"] = gen.PtrOf(gen.OneConstOf(
		PolicyContractProperties_Format_STATUS_Rawxml,
		PolicyContractProperties_Format_STATUS_RawxmlLink,
		PolicyContractProperties_Format_STATUS_Xml,
		PolicyContractProperties_Format_STATUS_XmlLink))
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
