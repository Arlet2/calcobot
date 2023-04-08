package model
import (
	"errors"
)

type Stack[V any] interface {
	Push(V)
	Pop() (V, error)
	Peek() (V, error)
	IsEmpty() bool
	PopAll() ([]V, error)
	Clear()
}

type Printable interface {
	Print()
}

type StackImpl[V any] struct {
	array []V
}

// Add some operation to top of the stack.
func (stack *StackImpl[V]) Push(value V) {
	newArray := make([]V, 0)

	newArray = append(newArray, value)
	newArray = append(newArray, stack.array...)

	stack.array = newArray
}

// Get some operation from top of the stack and remove it from the stack.
// Can produce error "Stack is empty" if stack is empty and you cannot get any element.
func (stack *StackImpl[V]) Pop() (value V, err error) {
	if stack.IsEmpty() {
		err = errors.New("stack is empty")
	} else {
		value = stack.array[0]
		stack.array = stack.array[1:]
	}
	
	return 
}

// Get some operation from top of the stack and DON'T remove it from the stack.
// Can produce error "Stack is empty" if stack is empty and you cannot get any element.
func (stack *StackImpl[V]) Peek() (value V, err error) {
	if stack.IsEmpty() {
		err = errors.New("stack is empty")
	} else {
		value = stack.array[0]
	}

	return
}

// Check is the stack is empty.
// Return true if stack hasn't got any element, otherwise true.
func (stack *StackImpl[V]) IsEmpty() bool {
	return len(stack.array) == 0 
}

// Remove all elements from the stack.
func (stack *StackImpl[V]) Clear() {
	stack.array = make([]V, 0)
}

// Get all elements from the stack and clear the stack.
// Can produce error "Stack is empty" if stack is empty and you cannot get any element.
func (stack *StackImpl[V]) PopAll() (values []V, err error) {
	if stack.IsEmpty() {
		err = errors.New("stack is empty")
	} else {
		values = stack.array
		stack.Clear()
	}

	return 
}