# shortest-path

This project derives shamelessly from the article [A Walkthrough of Dijkstra’s Algorithm (in JavaScript!)](https://medium.com/@adriennetjohnson/a-walkthrough-of-dijkstras-algorithm-in-javascript-e94b74192026) written by _**Adrienne Johnson**_.


The goal is to find the shortest path between 2 nodes in a given graph.

Let’s say we are at _Fullstack Academy_ in New York, and we want to know the shortest possible path to _Cafe Grumpy_ (see diagram). The weight of the edge between each node and its neighbors represents the distance it takes to walk from a node to the others.

The graph can be represented by a neighbors map. Each key in the neighbors list points to an array of edges extending from the node corresponding to the key. In a weighted graph, the neighbors list carries a second piece of information: the weight of each edge, or the cost of getting to that particular node.

<img src="./graph.jpeg">

in Go, a rough implementation may look like:
```go
type Edge struct {
  toNode  *Node
  weight  int
}

type Node struct {
  name string
  id    int
}

type Graph struct {
  nodes               map[int]*Node
  neighborsList       map[int][]Edge
}

```
To add a node to the graph, we push it into the collection of node values, which will help us iterate through them later, and we add a new entry in the neighbors list, setting its value to an empty array.

Create an empty graph
```go
func NewGraph() *Graph {
  neighbors := make(map[int][]Edge)
  nodes := make(map[int]*Node)
  return &Graph{nodes, neighbors}
}
```
Then from reading the graphdefinition.json file which contains the relationship between each node and their neighbors, we can build the graph using the buildGraph function.

The code is
```go
func main() {
	g := NewGraph()
	g.buildGraph("./graphdefinition.json")
	fmt.Println(g.findPathWithDijkstra(0, 5))
}
```
Here we use nodeId 0 and 5, which according to our graphdefinition represent respectively the "fullstack" and "Cafe Grumpy" nodes. We find the value **14** which corresponds to the shortest possible distance.

_buildGraph_ is in charge of creating the data structure that holds the graph data. It uses 2 loops where in the first loop, the nodes map is populated. In the second loop, the relationships between nodes are entered using the neighborsList map.  

## The Dijkstra Approach

- Move to a node that has not been visited, choosing the closest node to get to first.
- At that node, check how far each of its neighboring nodes are. Add the neighbor’s weight to the distance it took to get to the node we’re currently on. Keep in mind that we’re calculating the distance to reach those nodes before we visit them.
- Check whether that calculated distance is shorter than the previously known shortest distance to get to that node. If it is shorter, update our records to reflect the new shortest distance. We’ll also add this node to our line of nodes to visit next. That line will be arranged in order of shortest calculated distance to reach.

By calculating and continually updating the shortest distance to reach each node on the graph, the algorithm compiles the shortest path to the endpoint.

