package main

type StackItem struct {
	Value int
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack2 struct {
	items []*StackItem
	count int
}

type Stack struct {
	items []interface{}
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *StackItem) {
	s.items = append(s.items, n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() interface{} { //*Item {
	if s.count == 0 {
		return nil
	}
	defer func() {
		s.count--
		s.items = s.items[:s.count]
	}()
	return s.items[s.count-1]
}

// create a stack
func NewStack() *Stack {
	//	return &Stack{[]*Item{}, 0}
	return &Stack{}
}
