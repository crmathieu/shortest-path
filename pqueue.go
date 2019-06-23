package main

type Pqueue struct { 
    Items []*Node
} 

func NewPqueue() *Pqueue { 
    return &Pqueue{} 
} 

// insert item in queue based on priority (lower number is higher priority)
func (pq *Pqueue) Push(item *Node) {
    if len(pq.Items) > 0 { 
        for k, v := range(pq.Items) {
            if item.cumul < v.cumul  {
                w := append(pq.Items[:k], item)
                pq.Items = append(w, pq.Items[k:]...)
                return
            }
        }
    } 
    pq.Items = append(pq.Items, item)
} 
 
func (pq *Pqueue) isEmpty() bool {
    if len(pq.Items) == 0 {
        return true
    }
    return false
}

// remove 1st item in queue and return it
func (s *Pqueue) Dequeue() *Node { 
    defer func() { 
        s.Items = s.Items[1:] 
    }() 
    return s.Items[0] 
} 
 
