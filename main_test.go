package main

import (
	"testing"
)

func TestWeightedSum(t *testing.T) {
	tests := []struct {
		name     string
		a        []float64
		b        []float64
		expected float64
	}{
		{
			name:     "standard calculation",
			a:        []float64{8.5, 0.65, 1.2},
			b:        []float64{0.1, 0.2, 0},
			expected: 0.98,
		},
		{
			name:     "zeros",
			a:        []float64{0, 0, 0},
			b:        []float64{1, 2, 3},
			expected: 0,
		},
		{
			name:     "single element",
			a:        []float64{10},
			b:        []float64{0.5},
			expected: 5.0,
		},
	}

	const epsilon = 1e-9
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := weightedSum(tt.a, tt.b)
			if diff := got - tt.expected; diff < -epsilon || diff > epsilon {
				t.Errorf("weightedSum() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWeightedSumPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("weightedSum should have panicked due to unequal lengths")
		}
	}()
	weightedSum([]float64{1}, []float64{1, 2})
}
