// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	"encoding/json"
	v20231001s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
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

func Test_IstioServiceMesh_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from IstioServiceMesh to IstioServiceMesh via AssignProperties_To_IstioServiceMesh & AssignProperties_From_IstioServiceMesh returns original",
		prop.ForAll(RunPropertyAssignmentTestForIstioServiceMesh, IstioServiceMeshGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForIstioServiceMesh tests if a specific instance of IstioServiceMesh can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForIstioServiceMesh(subject IstioServiceMesh) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20231001s.IstioServiceMesh
	err := copied.AssignProperties_To_IstioServiceMesh(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual IstioServiceMesh
	err = actual.AssignProperties_From_IstioServiceMesh(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_IstioServiceMesh_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IstioServiceMesh via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIstioServiceMesh, IstioServiceMeshGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIstioServiceMesh runs a test to see if a specific instance of IstioServiceMesh round trips to JSON and back losslessly
func RunJSONSerializationTestForIstioServiceMesh(subject IstioServiceMesh) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IstioServiceMesh
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

// Generator of IstioServiceMesh instances for property testing - lazily instantiated by IstioServiceMeshGenerator()
var istioServiceMeshGenerator gopter.Gen

// IstioServiceMeshGenerator returns a generator of IstioServiceMesh instances for property testing.
func IstioServiceMeshGenerator() gopter.Gen {
	if istioServiceMeshGenerator != nil {
		return istioServiceMeshGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForIstioServiceMesh(generators)
	istioServiceMeshGenerator = gen.Struct(reflect.TypeOf(IstioServiceMesh{}), generators)

	return istioServiceMeshGenerator
}

// AddRelatedPropertyGeneratorsForIstioServiceMesh is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIstioServiceMesh(gens map[string]gopter.Gen) {
	gens["Components"] = gen.PtrOf(IstioComponentsGenerator())
}