package main

import (
	"fmt"
	"grokking/pkg/utils"
)

func neuralNetwork(input, weights []float64) float64 {
	prediction := utils.WeightedSum(input, weights)
	return prediction
}

func main() {
	weights := []float64{0.1, 0.2, 0}
	toes := []float64{8.5, 9.5, 9.9, 9.0}
	wlRecord := []float64{0.65, 0.8, 0.8, 0.9}
	numFans := []float64{1.2, 1.3, 0.5, 1.0}
	input := []float64{toes[0], wlRecord[0], numFans[0]}
	pred := neuralNetwork(input, weights)
	fmt.Printf("%.3f\n", pred)
}
