package main

type Queue struct { 
//    Items []interface{} 
    Items []*Node 
} 

func NewQueue() *Queue { 
    return &Queue{} 
} 
 
//func (s *Queue) Push(item interface{}) { 
func (s *Queue) Push(item *Node) { 
        s.Items = append(s.Items, item) 
} 
 
func (s *Queue) Prepend(item *Node) { 
    s.Items = append([]*Node{item}, s.Items...) 
} 

//func (s *Queue) Pop() interface{} { 
func (s *Queue) Pop() *Node { 
    defer func() { 
        s.Items = s.Items[1:] 
    }() 
    return s.Items[0] 
} 
 
func (s *Queue) isEmpty() bool { 
    if len(s.Items) == 0 { 
        return true 
    } 
    return false 
} 