package model

import (
	"fmt"
)

type NumberStack struct {
	StackImpl[float64]
}

func (stack NumberStack) Print() {
	if len(stack.array) == 0 {
		fmt.Println("EMPTY!")
		return
	}
	for _, value := range stack.array {
		fmt.Print(value, " <- ")
	}
	fmt.Println("bottom...")
}