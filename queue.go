package main

// Normal queue

type Queue struct {
	Items []*Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (s *Queue) isEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

// Push - append a new element in the queue
func (s *Queue) Push(item *Node) {
	s.Items = append(s.Items, item)
}

// Prepend - add a new element in the queue first position
func (s *Queue) Prepend(item *Node) {
	s.Items = append([]*Node{item}, s.Items...)
}

// DequeueFirst - removes 1st element of queue and returns it
func (s *Queue) DequeueFirst() *Node {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}
