package main

import "fmt"

func neutalNetwork(inout, weight float64) float64 {
	prediction := inout * weight
	return prediction
}

func main() {
	weight := 0.1
	input := []float64{8.5, 9.5, 10, 9}
	pred := neutalNetwork(input[0], weight)
	fmt.Printf("%.3f\n", pred)
}
