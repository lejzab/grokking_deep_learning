package neurals

import (
	"testing"
)

func TestMultipleInputsSingleOutput(t *testing.T) {
	weights := []float64{0.1, 0.2, 0}
	input := []float64{8.5, 0.65, 1.2}
	expected := 0.98

	got, err := MultipleInputsSingleOutput(input, weights)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	const epsilon = 1e-9
	if diff := got - expected; diff < -epsilon || diff > epsilon {
		t.Errorf("MultipleInputsSingleOutput() = %v, want %v", got, expected)
	}
}

func TestSingleInputMultipleOutputs(t *testing.T) {
	weights := []float64{0.3, 0.2, 0.9}
	input := 0.65
	expected := []float64{0.195, 0.13, 0.585}

	got := SingleInputMultipleOutputs(input, weights)

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

func TestMultipleInputsMultipleOutputs(t *testing.T) {
	weights := [][]float64{
		{0.1, 0.1, -0.3},
		{0.1, 0.2, 0.0},
		{0.0, 1.3, 0.1},
	}
	input := []float64{8.5, 0.65, 1.2}
	expected := []float64{0.555, 0.98, 0.965}

	got, err := MultipleInputsMultipleOutputs(input, weights)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(got) != len(expected) {
		t.Fatalf("MultipleInputsMultipleOutputs() returned %d elements, want %d", len(got), len(expected))
	}

	const epsilon = 1e-9
	for i := range got {
		if diff := got[i] - expected[i]; diff < -epsilon || diff > epsilon {
			t.Errorf("MultipleInputsMultipleOutputs() at index %d = %v, want %v", i, got[i], expected[i])
		}
	}
}
