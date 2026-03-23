package main

import (
	"fmt"
)

func neuralNetwork(input, weights []float64) float64 {
	prediction := weightedSum(input, weights)
	return prediction
}

func weightedSum(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("arrays must be of equal length")
	}
	sum := vectorSum(elementwiseMultiplication(a, b))
	return sum
}

func elementwiseMultiplication(a, b []float64) []float64 {
	if len(a) != len(b) {
		panic("arrays must be of equal length")
	}
	c := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] * b[i]
	}
	return c
}

func elementwiseAddition(a, b []float64) []float64 {
	if len(a) != len(b) {
		panic("arrays must be of equal length")
	}
	c := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

func vectorSum(a []float64) float64 {
	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return sum
}

func vectorAverage(a []float64) float64 {
	return vectorSum(a) / float64(len(a))
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
