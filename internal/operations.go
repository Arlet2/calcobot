package internal

import (
	"errors"
	"fmt"
)

var AllowedOperations = OperationDictImpl{
	[]operation{
		{'(', 0},
		{'+', 1},
		{'-', 1},
		{'*', 2},
		{'/', 2},
	},
}

type operation struct {
	symbol rune
	priority int
}

type OperationDict interface {
	IsOnDict(symbol rune) bool
	GetPriority(symbol rune) (int, error)
	GetOperation(symbol rune) (operation, error)
}

type OperationDictImpl struct{
	operations []operation
}

func (dict OperationDictImpl) IsOnDict(symbol rune) bool {
	for _, value := range dict.operations {
		if value.symbol == symbol {
			return true
		}
	}
	return false
}

func (dict OperationDictImpl) GetPriority(symbol rune) (int, error) {
	for _, value := range dict.operations {
		if value.symbol == symbol {
			return value.priority, nil
		}
	}

	return 0, errors.New(string(symbol)+" operation is not allowed")
}

func (dict OperationDictImpl) GetOperation(symbol rune) (operation, error) {
	for _, value := range dict.operations {
		if value.symbol == symbol {
			return value, nil
		}
	}

	return operation{}, errors.New(string(symbol)+" operation is not allowed")
}


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