package main

import (
	"fmt"
)

// findPathWithDijkstra - find the shortest path using the Dijsktra algorithm

func (g *Graph) findPathWithDijkstra(startNode, endNode *Node) string {

	distances := make(map[int]int)
	backtrace := make(map[int]int)

	pq := NewPqueue()

	distances[startNode.id] = 0

	for _, v := range g.nodes {
		if v.id != startNode.id {
			distances[v.id] = MAXINT
		}
	}

	// We add our starting node to the priority queue to get things kicked off
	pq.Insert(startNode, 0)

	// We access the first element in the queue and start checking its neighbors,
	// which we find using the neighbors list we made at the very beginning. We
	// add the neighbor’s weight to the time it took to get to the node we’re on.
	for !pq.isEmpty() {
		shortestStep := pq.DequeueFirst()
		currentNode := shortestStep.node

		for _, neighbor := range g.neighborsList[currentNode.id] {
			newDistance := distances[currentNode.id] + neighbor.weight

			// Then we check if the calculated time is less than the
			// time we currently have on file for this neighbor. If it is,
			// then we update our distances, we add this step to our backtrace,
			// and we add the neighbor to our priority queue!
			if newDistance < distances[neighbor.toNode.id] {
				distances[neighbor.toNode.id] = newDistance
				backtrace[neighbor.toNode.id] = currentNode.id
				pq.Insert(neighbor.toNode, newDistance)
			}
		}
	}

	// Once the end of our queue has been reached, all we have to do
	// is look through our backtrace to find the steps that will lead
	// us to the target node. We can look up target node in our 'distances'
	// object to find out just how long it will take to get there,
	// knowing that it’s the quickest route.

	path := NewQueue()
	path.Push(endNode)

	lastStep := endNode

	for lastStep.id != startNode.id {
		path.Prepend(g.nodes[backtrace[lastStep.id]])
		lastStep = g.nodes[backtrace[lastStep.id]]
	}

	// that's it
	output := ""
	for !path.isEmpty() {
		node := path.DequeueFirst()
		output = output + fmt.Sprintf("> %v ", node.name)
	}
	return fmt.Sprintf("Path is '%s' and distance is '%v'\n", output, distances[endNode.id])
}

func main() {
	fullstack = createNode("Fullstack")
	starbucks = createNode("Starbucks")
	insomnia = createNode("Insomnia Cookies")
	cafe = createNode("Cafe Grumpy")
	dig = createNode("Dig Inn")
	dubliner = createNode("Dubliner")

	g := NewGraph()
	g.BuildGraph()
	fmt.Println(g.findPathWithDijkstra(insomnia, dig))
}
