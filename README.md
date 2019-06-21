# shortest-path

This project derives shamelessly from the article [A Walkthrough of Dijkstra’s Algorithm (in JavaScript!)](https://medium.com/@adriennetjohnson/a-walkthrough-of-dijkstras-algorithm-in-javascript-e94b74192026) written by _**Adrienne Johnson**_.


The goal is to find the shortest path between 2 nodes in a given graph.

Let’s say we are at _Fullstack Academy_ in New York, and we want to know the shortest possible path to _Cafe Grumpy_ (see diagram). The weight of the edge between each node and its neighbors represents the time it takes to walk from a node to the others.

The graph can be represented by an adjacency list. Each node in the adjacency list points to an array of neighboring nodes, or in other words, the endpoint of every edge extending from that node. In a weighted graph, the adjacency list carries a second piece of information: the weight of each edge, or the cost of getting to that particular node.

<img src="./graph.jpeg">

in Go, a rough implementation may look like:
```go
type Edge struct {
    toNode  *Node
    weight  int
}

type Node struct {
    name string
    neighbors []Edge
}

type Graph struct {
    nodes []Node
    //adjacentList []Edge
}
```
To add a node to the graph, we push it into the collection of node values, which will help us iterate through them later, and we add a new entry in the adjacency list, setting its value to an empty array.

```go
  func (g *Graph) addNode(node string, edges Edge...) {
    this.nodes.push(node); 
    this.adjacencyList[node] = [];
  }
```


## Algorithm

Let the node at which we are starting be called the initial node. Let the distance of node Y be the distance from the initial node to Y. Dijkstra's algorithm will assign some initial distance values and will try to improve them step by step.

- Mark all nodes unvisited. Create a set of all the unvisited nodes called the unvisited set.
- Assign to every node a tentative distance value: set it to zero for our initial node and to infinity for all other nodes. Set the initial node as current.
- For the current node, consider all of its unvisited neighbours and calculate their tentative distances through the current node. 
- Compare the newly calculated tentative distance to the current assigned value and assign the smaller one. For example, if the current node A is marked with a distance of 6, and the edge connecting it with a neighbour B has length 2, then the distance to B through A will be 6 + 2 = 8. If B was previously marked with a distance greater than 8 then change it to 8. Otherwise, keep the current value.
- When we are done considering all of the unvisited neighbours of the current node, mark the current node as visited and remove it from the unvisited set. A visited node will never be checked again.
- If the destination node has been marked visited (when planning a route between two specific nodes) or if the smallest tentative distance among the nodes in the unvisited set is infinity (when planning a complete traversal; occurs when there is no connection between the initial node and remaining unvisited nodes), then stop. The algorithm has finished.
- Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it as the new "current node", and go back to step 3.

When planning a route, it is actually not necessary to wait until the destination node is "visited" as above: the algorithm can stop once the destination node has the smallest tentative distance among all "unvisited" nodes (and thus could be selected as the next "current").


