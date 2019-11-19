package stack

// Stack data structure
type Stack struct {
	top  *Element
	size int
}

// Element holds the value on the stack and points to next value on stack
type Element struct {
	value interface{}
	next  *Element
}

// New return instance of Stack
func New() Stack {
	return Stack{}
}

// Push value to top of the stack data structure
func (s *Stack) Push(value interface{}) interface{} {
	top := &Element{value: value, next: s.top}
	s.top = top
	s.size++
	return s.top.value
}

// Pop value from top of the stack data structure
func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		top := s.top

		newTop := s.top.next
		s.top = newTop

		s.size--
		return top.value
	}
	return nil
}

// Size returns the current length of stack data structure
func (s *Stack) Size() int {
	return s.size
}
