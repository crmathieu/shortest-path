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
  cnt                 int
  nodes               map[int]*Node
  neighborsList       map[int][]Edge
}

```
To add a node to the graph, we push it into the collection of node values, which will help us iterate through them later, and we add a new entry in the neighbors list, setting its value to an empty array.

Create a graph
```go
func NewGraph() *Graph {
  neighbors := make(map[int][]Edge)
  nodes := make(map[int]*Node)
  return &Graph{0, nodes, neighbors}
}
```

Create a Node
```go
func createNode(name string) *Node {
  return &Node{name, 0} //, MAXINT}
}
```
Add a node to the graph
```go
func (g *Graph) addNode(node *Node) {
  node.id = g.cnt
  g.cnt++
  g.nodes[node.id] = node
  g.neighborsList[node.id] = []Edge{}
}
```

To add edges to a node, we simply need to specify the weight between 2 nodes
```go
func (g *Graph) addEdge(node1, node2 *Node, weight int) {
  g.neighborsList[node1.id] = append(g.neighborsList[node1.id], Edge{node2, weight})
  g.neighborsList[node2.id] = append(g.neighborsList[node2.id], Edge{node1, weight})
}

```

## The Approach
In this simple example, it would be easy enough to scan the diagram and add some numbers to figure it out, but if I wanted to venture out of my neighborhood to a coffeeshop in, say, Midtown, the permutations of paths there would be much harder to calculate myself.

That’s where Dijkstra’s Algorithm comes in. Here is the basic idea of how it works:

- Move to a node that we haven’t visited, choosing the closest node to get to first.
- At that node, check how long it will take us to get to each of its neighboring nodes. Add the neighbor’s weight to the distance it took to get to the node we’re currently on. Keep in mind that we’re calculating the distance to reach those nodes before we visit them.
- Check whether that calculated distance is shorter than the previously known shortest distance to get to that node. If it is shorter, update our records to reflect the new shortest distance. We’ll also add this node to our line of nodes to visit next. That line will be arranged in order of shortest calculated distance to reach.

By calculating and continually updating the shortest distance to reach each node on the graph, the algorithm compiles the shortest path to the endpoint.

