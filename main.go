package main

import (
	"fmt"

	"github.com/rabingaire/sanskriti-vm/pkg/stack"
)

func main() {
	stack := stack.New()

	stack.Push("sanskriti-vm")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
}
