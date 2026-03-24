package neurals

import (
	"testing"
)

func TestMultipleInputsSingleOutput(t *testing.T) {
	weights := []float64{0.1, 0.2, 0}
	input := []float64{8.5, 0.65, 1.2}
	expected := 0.98

	got := MultipleInputsSingleOutput(input, weights)

	const epsilon = 1e-9
	if diff := got - expected; diff < -epsilon || diff > epsilon {
		t.Errorf("MultipleInputsSingleOutput() = %v, want %v", got, expected)
	}
}
