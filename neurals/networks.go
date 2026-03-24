package neurals

import "grokking/pkg/utils"

func MultipleInputsSingleOutput(input, weights []float64) float64 {
	prediction := utils.WeightedSum(input, weights)
	return prediction
}

func SingleInputMultipleOutputs(input float64, weights []float64) []float64 {
	return utils.ElementwiseMultiplicationScalar(input, weights)
}
