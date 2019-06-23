package main

type Pqueue struct { 
    Items []Node
} 

func NewQueue() *Pqueue { 
    return &Pqueue{} 
} 
 
func (s *Pqueue) Push(item Node) {
    if len(s.Items) == 0 { 
        s.Items = append(s.Items, item)
    } else {
        for k, v := range(s.Items) {
            if v.cumul > item.cumul {
                s.Items = s.Items[:k]
            }
        }
    } 
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