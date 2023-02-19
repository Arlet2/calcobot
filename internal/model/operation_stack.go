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

// Print all elements from the stack.
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

// Add some operation to top of the stack.
func (stack *OperationStackImpl) Push(value operation) {
	newArray := make([]operation, 0)

	newArray = append(newArray, value)
	newArray = append(newArray, stack.array...)

	stack.array = newArray
}

// Get some operation from top of the stack and remove it from the stack.
// Can produce error "Stack is empty" if stack is empty and you cannot get any element.
func (stack *OperationStackImpl) Pop() (operation, error) {
	if stack.IsEmpty() {
		return operation{}, errors.New("stack is empty")
	}
	value := stack.array[0]

	stack.array = stack.array[1:]

	return value, nil
}

// Get some operation from top of the stack and DON'T remove it from the stack.
// Can produce error "Stack is empty" if stack is empty and you cannot get any element.
func (stack *OperationStackImpl) Peek() (operation, error) {
	if stack.IsEmpty() {
		return operation{}, errors.New("stack is empty")
	}

	return stack.array[0], nil
}

// Check is the stack is empty.
// Return true if stack hasn't got any element, otherwise true.
func (stack *OperationStackImpl) IsEmpty() bool {
	return len(stack.array) == 0 
}

// Remove all elements from the stack.
func (stack *OperationStackImpl) Clear() {
	stack.array = make([]operation, 0)
}

// Get all elements from the stack and clear the stack.
// Can produce error "Stack is empty" if stack is empty and you cannot get any element.
func (stack *OperationStackImpl) PopAll() ([]operation, error) {
	if stack.IsEmpty() {
		return []operation{}, errors.New("stack is empty")
	}

	array := stack.array

	stack.Clear()

	return array, nil
}