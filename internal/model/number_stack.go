package model

import (
	"errors"
	"fmt"
)

type NumberStack interface {
	Push(value float64)
	Pop() (float64, error)
}

type NumberStackImpl struct {
	array []float64
}

func (stack NumberStackImpl) Print() {
	if len(stack.array) == 0 {
		fmt.Println("EMPTY!")
		return
	}
	for _, value := range stack.array {
		fmt.Print(value, " <- ")
	}
	fmt.Println("bottom...")
}

func (stack *NumberStackImpl) Push(value float64)  {
	newArray := make([]float64, 0)

	newArray = append(newArray, value)
	newArray = append(newArray, stack.array...)

	stack.array = newArray

}

func (stack *NumberStackImpl) Pop() (float64, error) {
	if len(stack.array) == 0 {
		return 0, errors.New("stack is empty")
	}
	value := stack.array[0]

	stack.array = stack.array[1:]

	return value, nil
}