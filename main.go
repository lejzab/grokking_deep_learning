package main

import (
	"fmt"
	"grokking/neurals"
)

func main() {
	weights := []float64{0.1, 0.2, 0}
	toes := []float64{8.5, 9.5, 9.9, 9.0}
	wlRecord := []float64{0.65, 0.8, 0.8, 0.9}
	numFans := []float64{1.2, 1.3, 0.5, 1.0}
	input := []float64{toes[0], wlRecord[0], numFans[0]}
	pred := neurals.MultipleInputsSingleOutput(input, weights)
	fmt.Printf("%.3f\n", pred)

	weights = []float64{0.3, 0.2, 0.9}
	wlRecord = []float64{0.65, 0.8, 0.8, 0.9}
	inputSingle := wlRecord[0]
	predSingle := neurals.SingleInputMultipleOutputs(inputSingle, weights)
	fmt.Printf("%.3f\n", predSingle)
}
