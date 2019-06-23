package main
import (
	"fmt"
)

func (g *Graph) findPathWithDijkstra(startNode, endNode *Node) string {
	times := make(map[int]int)

	backtrace := make(map[int]int)
	pq := NewPqueue();

	times[startNode.id] = 0;
  
	for _, v := range(g.nodes) {
		if v.id != startNode.id {
			times[v.id] = MAXINT
		}
	}

	// We add our starting node to the priority queue to get things kicked off
	startNode.cumul = 0
	pq.Push(startNode)

	for ; !pq.isEmpty()	; {
		shortestStep := pq.Dequeue()
		currentNode := shortestStep

		for _, neighbor := range g.adjacentList[currentNode.name] {
			vtime := times[currentNode.id] + neighbor.weight

			if (vtime < times[neighbor.toNode.id]) {
				times[neighbor.toNode.id] = vtime;
				backtrace[neighbor.toNode.id] = currentNode.id;
				neighbor.toNode.cumul = vtime
				pq.Push(neighbor.toNode);
			}
		}
	}	  

	path := NewQueue()
	path.Push(endNode)

  	lastStep := endNode
	  
  	for ;lastStep.id != startNode.id; {
		path.Prepend(g.nodes[backtrace[lastStep.id]])
		lastStep = g.nodes[backtrace[lastStep.id]]
	}
	output := ""
	for ;!path.isEmpty(); {
		node := path.Pop()
		output = output + fmt.Sprintf(">(%v, %v)", node.name, node.cumul)
	}  
  	return fmt.Sprintf("Path is '%s' and time is '%v'\n", output, times[endNode.id])
}


func main() {
	fullstack 	= createNode("Fullstack")
	starbucks 	= createNode("Starbucks")
	insomnia 	= createNode("Insomnia Cookies")
	cafe 		= createNode("Cafe Grumpy")
	dig 		= createNode("Dig Inn")
	dubliner 	= createNode("Dubliner")

	g := NewGraph()
	g.Build()
	fmt.Println(g.findPathWithDijkstra(dubliner, cafe))
}
