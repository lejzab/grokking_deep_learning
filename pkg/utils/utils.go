package utils

func WeightedSum(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("arrays must be of equal length")
	}
	sum := VectorSum(ElementwiseMultiplication(a, b))
	return sum
}

func ElementwiseMultiplication(a, b []float64) []float64 {
	if len(a) != len(b) {
		panic("arrays must be of equal length")
	}
	c := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] * b[i]
	}
	return c
}

func ElementwiseMultiplicationScalar(input float64, weights []float64) []float64 {
	c := make([]float64, len(weights))
	for i := 0; i < len(weights); i++ {
		c[i] = weights[i] * input
	}
	return c
}

func ElementwiseAddition(a, b []float64) []float64 {
	if len(a) != len(b) {
		panic("arrays must be of equal length")
	}
	c := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

func VectorSum(a []float64) float64 {
	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return sum
}

func VectorAverage(a []float64) float64 {
	return VectorSum(a) / float64(len(a))
}
