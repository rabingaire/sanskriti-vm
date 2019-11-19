package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output interface{}
		Error  error
	}{
		{
			Name:   "Push one string to stack",
			Input:  []string{"hello"},
			Output: "hello",
		},
		{
			Name:   "Empty stack",
			Input:  []string{},
			Output: nil,
		},
		{
			Name:   "Push three string to stack",
			Input:  []string{"hello", "world", "go"},
			Output: "go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			s := New()

			var output interface{}
			for _, i := range tt.Input {
				output = s.Push(i)
			}

			if !reflect.DeepEqual(output, tt.Output) {
				t.Fatalf("Push() error:\ngot %v\nwant %v", output, tt.Output)
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output interface{}
		Error  error
	}{
		{
			Name:   "Pop one string from stack",
			Input:  []string{"hello"},
			Output: "hello",
		},
		{
			Name:   "Empty stack",
			Input:  []string{},
			Output: nil,
		},
		{
			Name:   "Push three string to stack",
			Input:  []string{"hello", "world", "go"},
			Output: "go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			s := New()

			for _, i := range tt.Input {
				s.Push(i)
			}

			output := s.Pop()
			if !reflect.DeepEqual(output, tt.Output) {
				t.Fatalf("Pop() error:\ngot %v\nwant %v", output, tt.Output)
			}
		})
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output int
		Error  error
	}{
		{
			Name:   "Pop one string from stack",
			Input:  []string{"hello"},
			Output: 1,
		},
		{
			Name:   "Empty stack",
			Input:  []string{},
			Output: 0,
		},
		{
			Name:   "Push three string to stack",
			Input:  []string{"hello", "world", "go"},
			Output: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			s := New()

			for _, i := range tt.Input {
				s.Push(i)
			}

			output := s.Size()
			if !reflect.DeepEqual(output, tt.Output) {
				t.Fatalf("Size() error:\ngot %v\nwant %v", output, tt.Output)
			}
		})
	}
}
