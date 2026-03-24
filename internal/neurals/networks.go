package neurals

import "grokking/pkg/math"

// MultipleInputsSingleOutput to sieć neuronowa z wieloma wejściami i jednym wyjściem.
// Zwraca błąd, jeśli dane wejściowe i wagi mają różne długości.
func MultipleInputsSingleOutput(input, weights []float64) (float64, error) {
	prediction, err := math.WeightedSum(input, weights)
	if err != nil {
		return 0, err
	}
	return prediction, nil
}

// SingleInputMultipleOutputs to sieć neuronowa z jednym wejściem i wieloma wyjściami.
func SingleInputMultipleOutputs(input float64, weights []float64) []float64 {
	return math.ElementwiseMultiplicationScalar(input, weights)
}

// MultipleInputsMultipleOutputs to sieć neuronowa z wieloma wejściami i wieloma wyjściami.
// Zwraca błąd, jeśli wymiary wejścia i macierzy wag są niezgodne.
func MultipleInputsMultipleOutputs(inputs []float64, weights [][]float64) ([]float64, error) {
	pred, err := math.VectorMatrixMultiplication(inputs, weights)
	if err != nil {
		return nil, err
	}
	return pred, nil
}
