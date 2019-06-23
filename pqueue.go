package main

// priority queue

type Pqueue struct { 
    Items []*Node
} 

func NewPqueue() *Pqueue { 
    return &Pqueue{} 
} 

func (pq *Pqueue) isEmpty() bool {
    if len(pq.Items) == 0 {
        return true
    }
    return false
}

// Insert - place a new item in queue based on priority (lower number is higher priority)
func (pq *Pqueue) Insert(item *Node) {
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
 

// remove 1st item in queue and return it
func (s *Pqueue) DequeueFirst() *Node { 
    defer func() { 
        s.Items = s.Items[1:] 
    }() 
    return s.Items[0] 
} 
 
