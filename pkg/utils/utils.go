package utils

// WeightedSum oblicza sumę ważoną dwóch wektorów (iloczyn skalarny).
// Oczekuje, że oba wektory mają taką samą długość, w przeciwnym razie wywołuje panic.
func WeightedSum(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("arrays must be of equal length")
	}
	sum := VectorSum(ElementwiseMultiplication(a, b))
	return sum
}

// ElementwiseMultiplication wykonuje mnożenie element po elemencie (iloczyn Hadamarda).
// Oczekuje, że oba wektory mają taką samą długość, w przeciwnym razie wywołuje panic.
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

// ElementwiseMultiplicationScalar mnoży każdy element wektora przez skalar.
func ElementwiseMultiplicationScalar(input float64, weights []float64) []float64 {
	c := make([]float64, len(weights))
	for i := 0; i < len(weights); i++ {
		c[i] = weights[i] * input
	}
	return c
}

// ElementwiseAddition dodaje do siebie dwa wektory element po elemencie.
// Oczekuje, że oba wektory mają taką samą długość, w przeciwnym razie wywołuje panic.
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

// VectorSum sumuje wszystkie elementy w wektorze.
func VectorSum(a []float64) float64 {
	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return sum
}

// VectorAverage oblicza średnią arytmetyczną elementów wektora.
// Jeśli wektor jest pusty, zwraca 0.
func VectorAverage(a []float64) float64 {
	if len(a) == 0 {
		return 0
	}
	return VectorSum(a) / float64(len(a))
}

// VectorMatrixMultiplication wykonuje mnożenie wektora przez macierz.
// Oczekuje, że długość wektora zgadza się z liczbą kolumn macierzy, w przeciwnym razie wywołuje panic.
// Jeśli macierz jest pusta, zwraca pusty wektor.
func VectorMatrixMultiplication(vector []float64, matrix [][]float64) []float64 {
	if len(matrix) == 0 {
		return []float64{}
	}
	if len(vector) != len(matrix[0]) {
		panic("vector and matrix must have the same length")
	}
	result := make([]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		result[i] = WeightedSum(vector, matrix[i])
	}
	return result
}
