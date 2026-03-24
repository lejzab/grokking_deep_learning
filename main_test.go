package main

import (
	"grokking/neurals"
	"testing"
)

func TestIntegration(t *testing.T) {
	weights := []float64{0.1, 0.2, 0}
	input := []float64{8.5, 0.65, 1.2}
	expected := 0.98

	got := neurals.MultipleInputsSingleOutput(input, weights)

	const epsilon = 1e-9
	if diff := got - expected; diff < -epsilon || diff > epsilon {
		t.Errorf("MultipleInputsSingleOutput() = %v, want %v", got, expected)
	}
}

func TestIntegrationSingleInputMultipleOutputs(t *testing.T) {
	weights := []float64{0.3, 0.2, 0.9}
	input := 0.65
	expected := []float64{0.195, 0.13, 0.585}

	got := neurals.SingleInputMultipleOutputs(input, weights)

	if len(got) != len(expected) {
		t.Fatalf("SingleInputMultipleOutputs() returned %d elements, want %d", len(got), len(expected))
	}

	const epsilon = 1e-9
	for i := range got {
		if diff := got[i] - expected[i]; diff < -epsilon || diff > epsilon {
			t.Errorf("SingleInputMultipleOutputs() at index %d = %v, want %v", i, got[i], expected[i])
		}
	}
}
