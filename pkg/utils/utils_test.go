package utils

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
			got := WeightedSum(tt.a, tt.b)
			if diff := got - tt.expected; diff < -epsilon || diff > epsilon {
				t.Errorf("WeightedSum() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWeightedSumPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("WeightedSum should have panicked due to unequal lengths")
		}
	}()
	WeightedSum([]float64{1}, []float64{1, 2})
}

func TestElementwiseMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		a        []float64
		b        []float64
		expected []float64
	}{
		{
			name:     "standard",
			a:        []float64{1, 2, 3},
			b:        []float64{4, 5, 6},
			expected: []float64{4, 10, 18},
		},
		{
			name:     "with zero",
			a:        []float64{1, 0, 3},
			b:        []float64{4, 5, 6},
			expected: []float64{4, 0, 18},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ElementwiseMultiplication(tt.a, tt.b)
			if len(got) != len(tt.expected) {
				t.Fatalf("length mismatch: got %d, want %d", len(got), len(tt.expected))
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("at index %d: got %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestElementwiseMultiplicationScalar(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		weights  []float64
		expected []float64
	}{
		{
			name:     "standard",
			input:    0.5,
			weights:  []float64{1, 2, 3},
			expected: []float64{0.5, 1, 1.5},
		},
		{
			name:     "with zero scalar",
			input:    0,
			weights:  []float64{1, 2, 3},
			expected: []float64{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ElementwiseMultiplicationScalar(tt.input, tt.weights)
			if len(got) != len(tt.expected) {
				t.Fatalf("length mismatch: got %d, want %d", len(got), len(tt.expected))
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("at index %d: got %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestElementwiseAddition(t *testing.T) {
	tests := []struct {
		name     string
		a        []float64
		b        []float64
		expected []float64
	}{
		{
			name:     "standard",
			a:        []float64{1, 2, 3},
			b:        []float64{4, 5, 6},
			expected: []float64{5, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ElementwiseAddition(tt.a, tt.b)
			if len(got) != len(tt.expected) {
				t.Fatalf("length mismatch: got %d, want %d", len(got), len(tt.expected))
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("at index %d: got %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestVectorSum(t *testing.T) {
	tests := []struct {
		name     string
		a        []float64
		expected float64
	}{
		{
			name:     "standard",
			a:        []float64{1, 2, 3},
			expected: 6,
		},
		{
			name:     "empty",
			a:        []float64{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VectorSum(tt.a)
			if got != tt.expected {
				t.Errorf("VectorSum() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorAverage(t *testing.T) {
	tests := []struct {
		name     string
		a        []float64
		expected float64
	}{
		{
			name:     "standard",
			a:        []float64{1, 2, 3},
			expected: 2,
		},
		{
			name:     "single",
			a:        []float64{5},
			expected: 5,
		},
		{
			name:     "empty",
			a:        []float64{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VectorAverage(tt.a)
			if got != tt.expected {
				t.Errorf("VectorAverage() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorMatrixMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		vector   []float64
		matrix   [][]float64
		expected []float64
	}{
		{
			name:   "standard 3x3",
			vector: []float64{8.5, 0.65, 1.2},
			matrix: [][]float64{
				{0.1, 0.1, -0.3},
				{0.1, 0.2, 0.0},
				{0.0, 1.3, 0.1},
			},
			expected: []float64{0.555, 0.98, 0.965},
		},
		{
			name:     "empty matrix",
			vector:   []float64{1, 2, 3},
			matrix:   [][]float64{},
			expected: []float64{},
		},
	}

	const epsilon = 1e-9
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VectorMatrixMultiplication(tt.vector, tt.matrix)
			if len(got) != len(tt.expected) {
				t.Fatalf("length mismatch: got %d, want %d", len(got), len(tt.expected))
			}
			for i := range got {
				if diff := got[i] - tt.expected[i]; diff < -epsilon || diff > epsilon {
					t.Errorf("at index %d: got %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}
