package math

import "fmt"

// WeightedSum oblicza sumę ważoną dwóch wektorów (iloczyn skalarny).
// Zwraca błąd, jeśli wektory mają różne długości.
func WeightedSum(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("wektory muszą mieć taką samą długość: %d != %d", len(a), len(b))
	}
	res, err := ElementwiseMultiplication(a, b)
	if err != nil {
		return 0, err
	}
	sum := VectorSum(res)
	return sum, nil
}

// ElementwiseMultiplication wykonuje mnożenie element po elemencie (iloczyn Hadamarda).
// Zwraca błąd, jeśli wektory mają różne długości.
func ElementwiseMultiplication(a, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("wektory muszą mieć taką samą długość: %d != %d", len(a), len(b))
	}
	c := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] * b[i]
	}
	return c, nil
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
// Zwraca błąd, jeśli wektory mają różne długości.
func ElementwiseAddition(a, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("wektory muszą mieć taką samą długość: %d != %d", len(a), len(b))
	}
	c := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	return c, nil
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
// Zwraca błąd, jeśli długość wektora nie zgadza się z liczbą kolumn macierzy
// lub jeśli macierz jest pusta.
func VectorMatrixMultiplication(vector []float64, matrix [][]float64) ([]float64, error) {
	if len(matrix) == 0 {
		return nil, fmt.Errorf("macierz wag nie może być pusta")
	}
	if len(vector) != len(matrix[0]) {
		return nil, fmt.Errorf("wektor i macierz muszą mieć zgodne wymiary: len(vector)=%d, len(matrix[0])=%d", len(vector), len(matrix[0]))
	}
	result := make([]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		val, err := WeightedSum(vector, matrix[i])
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}
