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

func pow(a float64, b float64) (result float64) {
	result = 1

	if b > 0 {
		for i := 0; i < int(b); i++ {
			result *= a
		}
	} else if b < 0{
		for i := 0; i < -int(b); i++ {
			result /= a
		}
	}

	return
}

func unaryMinus(a float64, _ float64) float64 {
	return -a
}

func factorial(a float64, _ float64) (value float64) {
	if a == 0 || a == 1 {
		value = 1
	} else if a > 34 {
		value = -1
	} else {
		value = 2
		for i := 3; i <= int(a); i++ {
			value *= float64(i)
		}
	}
	return
}