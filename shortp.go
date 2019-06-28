package main

import (
	"fmt"
)

// findPathWithDijkstra - find the shortest path using the Dijsktra algorithm

func (g *Graph) findPathWithDijkstra(startNodeID, endNodeID int) (string, int) {

	distances := make(map[int]int)
	backtrace := make(map[int]int)
	visited := make(map[int]int)
	pq := NewPqueue()

	// initialize the distance of the source node as 0
	distances[startNodeID] = 0

	// and the other distances to MAXINT
	for _, v := range g.nodes {
		if v.id != startNodeID {
			distances[v.id] = MAXINT
		}
	}

	// We add our starting node to the priority queue to get things kicked off
	pq.Insert(g.nodes[startNodeID], 0)

	// We access the first element in the queue and start checking its neighbors,
	// which we find using the neighbors list we made at the very beginning. We
	// add the neighbor’s weight to the time it took to get to the node we’re on.
	for !pq.isEmpty() {
		shortestStep := pq.DequeueFirst()
		currentNode := shortestStep.node
		visited[currentNode.id] = 1

		for _, neighbor := range g.neighborsList[currentNode.id] {
			if _, ok := visited[neighbor.toNode.id]; !ok {
				newDistance := distances[currentNode.id] + neighbor.weight

				// Then we check if the calculated distance is less than the
				// distance we currently have on file for this neighbor. If it is,
				// then we update our distances, we add this step to our backtrace,
				// and we add the neighbor to our priority queue!
				if newDistance < distances[neighbor.toNode.id] {
					distances[neighbor.toNode.id] = newDistance
					backtrace[neighbor.toNode.id] = currentNode.id
					pq.Insert(neighbor.toNode, newDistance)
				}
			}
		}
	}

	// Once the end of our queue has been reached, all we have to do
	// is look through our backtrace map to find the steps that lead to
	// the target node. We can look up target node in our 'distances'
	// map to find out just how long it will take to get there, knowing
	// that it’s the quickest route.

	lastStep := g.nodes[endNodeID]
	output := ""

	for lastStep.id != startNodeID {
		output = fmt.Sprintf("> %v ", lastStep.name) + output
		lastStep = g.nodes[backtrace[lastStep.id]]
	}
	output = fmt.Sprintf("%v ", g.nodes[startNodeID].name) + output

	// that's it
	return output, distances[endNodeID]
}

func main() {
	g := NewGraph()
	g.buildGraph("./graphdefinition.json")
	path, distance := g.findPathWithDijkstra(0, 5)
	fmt.Printf("Shortest Path from '%s' to '%s' is '%s' and distance is '%v'\n", g.nodes[0].name, g.nodes[5].name, path, distance)
}
