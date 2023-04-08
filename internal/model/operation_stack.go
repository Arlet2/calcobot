package model

import (
	"fmt"
)


type OperationStack struct {
	StackImpl[operation]
}

// Print all elements from the stack.
func (stack OperationStack) Print() {
	if stack.IsEmpty() {
		fmt.Println("EMPTY!")
		return
	}
	for _, value := range stack.array {
		fmt.Print(string(value.symbol)+" <- ")
	}
	fmt.Println("bottom...")
}