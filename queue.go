package main

type Queue struct { 
    Items []interface{} 
} 

func NewQueue() *Queue { 
    return &Queue{} 
} 
 
func (s *Queue) Push(item interface{}) { 
    s.Items = append(s.Items, item) 
} 
 
func (s *Queue) Pop() interface{} { 
    defer func() { 
        s.Items = s.Items[1:] 
    }() 
    return s.Items[0] 
} 
 
func (s *Queue) IsEmpty() bool { 
    if len(s.Items) == 0 { 
        return true 
    } 
    return false 
} 