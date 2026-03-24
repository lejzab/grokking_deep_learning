package neurals

import "grokking/pkg/utils"

// MultipleInputsSingleOutput to sieć neuronowa z wieloma wejściami i jednym wyjściem.
func MultipleInputsSingleOutput(input, weights []float64) float64 {
	prediction := utils.WeightedSum(input, weights)
	return prediction
}

// SingleInputMultipleOutputs to sieć neuronowa z jednym wejściem i wieloma wyjściami.
func SingleInputMultipleOutputs(input float64, weights []float64) []float64 {
	return utils.ElementwiseMultiplicationScalar(input, weights)
}

// MultipleInputsMultipleOutputs to sieć neuronowa z wieloma wejściami i wieloma wyjściami.
func MultipleInputsMultipleOutputs(inputs []float64, weights [][]float64) []float64 {
	pred := utils.VectorMatrixMultiplication(inputs, weights)
	return pred
}
