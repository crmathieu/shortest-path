package main

const MAXINT = 2147483647 // 0x7fffffff

type Node struct {
	name string
	id   int
}

type Edge struct {
	toNode *Node
	weight int
}

type Graph struct {
	cnt                int
	nodes              map[int]*Node
	neighborsList      map[int][]Edge
}

func NewGraph() *Graph {
	neighbors := make(map[int][]Edge)
	nodes := make(map[int]*Node)
	return &Graph{0, nodes, neighbors}
}

func createNode(name string) *Node {
	return &Node{name, 0}
}

func (g *Graph) addNode(node *Node) {
	node.id = g.cnt
	g.cnt++
	g.nodes[node.id] = node
	g.neighborsList[node.id] = []Edge{}
}

func (g *Graph) addEdge(node1, node2 *Node, weight int) {
	g.neighborsList[node1.id] = append(g.neighborsList[node1.id], Edge{node2, weight})
	g.neighborsList[node2.id] = append(g.neighborsList[node2.id], Edge{node1, weight})
}

var fullstack, starbucks, insomnia, cafe, dig, dubliner *Node

func (g *Graph) BuildGraph() {

	g.addNode(fullstack)
	g.addNode(starbucks)
	g.addNode(insomnia)
	g.addNode(cafe)
	g.addNode(dig)
	g.addNode(dubliner)

	g.addEdge(fullstack, starbucks, 6)
	g.addEdge(fullstack, dig, 7)
	g.addEdge(fullstack, dubliner, 2)
	g.addEdge(dig, dubliner, 4)
	g.addEdge(dig, cafe, 9)
	g.addEdge(cafe, insomnia, 5)
	g.addEdge(insomnia, starbucks, 6)
	g.addEdge(insomnia, dubliner, 7)
	g.addEdge(starbucks, dubliner, 3)
}
