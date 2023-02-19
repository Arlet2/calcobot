package model

var AllowedOperations = OperationDictImpl{
	[]operation{
		{'(', 0, false, nil},
		{'+', 1, true, add},
		{'-', 1, true, sub},
		{'*', 2, true, mul},
		{'/', 2, true, div},
		{'^', 3, true, pow},
		{'~', 4, false, unaryMinus}, // унарный минус
		{'!', 4, false, factorial}, // факториал
	},
}

type operation struct {
	symbol rune
	priority int
	isBinary bool
	calc func (a float64, b float64) float64
}

func add(a float64, b float64) float64 {
	return a+b
}

func sub(a float64, b float64) float64 {
	return a-b
}

func mul(a float64, b float64) float64 {
	return a*b
}

func div(a float64, b float64) float64 {
	return a/b
}

func pow(a float64, b float64) float64 {
	var result float64
	result = 1

	for i := 0; i < int(b); i++ {
		result *= a
	}

	return result
}

func unaryMinus(a float64, _ float64) float64 {
	return -a
}

// добавить многопоточку
func factorial(a float64, _ float64) float64 {
	result := 2
	for i := 3; i <= int(a); i++ {
		result *= i
	}

	return float64(result)
}