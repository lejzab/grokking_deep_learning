package neurals

import "grokking/pkg/utils"

func MultipleInputsSingleOutput(input, weights []float64) float64 {
	prediction := utils.WeightedSum(input, weights)
	return prediction
}
