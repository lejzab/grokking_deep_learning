package main

import (
	"grokking_neural_nets/pkg/utils"
	"testing"
)

func TestNeuralNetwork(t *testing.T) {
	weights := []float64{0.1, 0.2, 0}
	input := []float64{8.5, 0.65, 1.2}
	expected := 0.98

	got := neuralNetwork(input, weights)

	const epsilon = 1e-9
	if diff := got - expected; diff < -epsilon || diff > epsilon {
		t.Errorf("neuralNetwork() = %v, want %v", got, expected)
	}
}

func TestWeightedSum(t *testing.T) {
	// Ten test teraz głównie sprawdza, czy integracja z utils działa
	a := []float64{8.5, 0.65, 1.2}
	b := []float64{0.1, 0.2, 0}
	expected := 0.98

	got := utils.WeightedSum(a, b)

	const epsilon = 1e-9
	if diff := got - expected; diff < -epsilon || diff > epsilon {
		t.Errorf("utils.WeightedSum() = %v, want %v", got, expected)
	}
}
