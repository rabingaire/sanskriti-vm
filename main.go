package main

import (
	"fmt"

	"github.com/rabingaire/sanskriti-vm/pkg/stack"
)

// TODO: pass instruction sets from outer world
// refactor main.go
// handle case for invalid instruction set

// ByteCode ...
type ByteCode int

// ByteCodes
const (
	ICONST ByteCode = iota
	IADD
	PRINT
	HALT
)

// Machine ...
type Machine struct {
	stack              stack.Stack
	instructions       []ByteCode
	instructionPointer int
}

// Instruction Set of:
// (print (+ 1 (+ 2 3)))

// New ...
func New() *Machine {
	return &Machine{
		stack: stack.New(),
		instructions: []ByteCode{
			ICONST, 1,
			ICONST, 2,
			ICONST, 3,
			IADD,
			IADD,
			PRINT,
			HALT,
		},
	}
}

func main() {
	machine := New()

	for machine.instructionPointer < len(machine.instructions) {
		switch machine.instructions[machine.instructionPointer] {
		case ICONST:
			machine.instructionPointer++
			machine.stack.Push(machine.instructions[machine.instructionPointer])
		case IADD:
			right := machine.stack.Pop().(ByteCode)
			left := machine.stack.Pop().(ByteCode)
			machine.stack.Push(left + right)
		case PRINT:
			result := machine.stack.Pop()
			fmt.Println(result)
		case HALT:
			return
		}
		machine.instructionPointer++
	}
}
