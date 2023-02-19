package model

import (
	"errors"
	"fmt"
)

type OperationStack interface {
	Push(operation)
	Pop() (operation, error)
	Peek() (operation, error)
	IsEmpty() bool
	PopAll() ([]operation, error)
	Clear()
}

type OperationStackImpl struct {
	array []operation
}

func (stack OperationStackImpl) Print() {
	if stack.IsEmpty() {
		fmt.Println("EMPTY!")
		return
	}
	for _, value := range stack.array {
		fmt.Print(string(value.symbol)+" <- ")
	}
	fmt.Println("bottom...")
}

func (stack *OperationStackImpl) Push(value operation) {
	newArray := make([]operation, 0)

	newArray = append(newArray, value)
	newArray = append(newArray, stack.array...)

	stack.array = newArray
}

func (stack *OperationStackImpl) Pop() (operation, error) {
	if stack.IsEmpty() {
		return operation{}, errors.New("stack is empty")
	}
	value := stack.array[0]

	stack.array = stack.array[1:]

	return value, nil
}

func (stack *OperationStackImpl) Peek() (operation, error) {
	if stack.IsEmpty() {
		return operation{}, errors.New("stack is empty")
	}

	return stack.array[0], nil
}

func (stack *OperationStackImpl) IsEmpty() bool {
	return len(stack.array) == 0 
}

func (stack *OperationStackImpl) Clear() {
	stack.array = make([]operation, 0)
}

func (stack *OperationStackImpl) PopAll() ([]operation, error) {
	if stack.IsEmpty() {
		return []operation{}, errors.New("stack is empty")
	}

	array := stack.array

	stack.Clear()

	return array, nil
}