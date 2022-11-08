package stack_string

type node struct {
	value *string
	lower *node
}

type stack struct {
	top *node
}

// Makes new stack (data type: string).
func NewStackString() *stack {
	return &stack{}
}

// Create new element on the top of the stack.
// Stack elements number rises by one.
// TC: O(1)
func (s *stack) Push(value *string) {
	s.top = &node{value: value, lower: s.top}
}

// Gets value and removes element from the top of the stack.
// Stack elements number reduces by one.
// TC: O(1)
func (s *stack) Pop() *string {
	if s.top == nil {
		return nil
	}

	value := s.top.value

	s.top = s.top.lower

	return value
}
