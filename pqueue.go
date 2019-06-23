package main

// priority queue

type Item struct {
	node  *Node
	cumul int
}

type Pqueue struct {
	Items []Item
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
func (pq *Pqueue) Insert(n *Node, distance int) {
	if len(pq.Items) > 0 {
		for k, v := range pq.Items {
			if distance < v.cumul {
				w := append(pq.Items[:k], Item{n, distance})
				pq.Items = append(w, pq.Items[k:]...)
				return
			}
		}
	}
	pq.Items = append(pq.Items, Item{n, distance})
}

// remove 1st item in queue and return it
func (s *Pqueue) DequeueFirst() Item {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}
