package main

type Node struct {
	name string
	id    int
	cumul int
}

type Edge struct {
    toNode  *Node
    weight  int
}

func NewGraph() *Graph {
	adj := make(map[string][]Edge)
	dist := make(map[string]int)
	nodes := make(map[int]*Node)
//	return &Graph{&Node{}, &Node{}, []*Node{}, adj}
	return &Graph{&Node{}, &Node{}, 0, nodes, adj, dist}
}

type Graph struct {
	startNode, endNode  		*Node
	cnt 			int
//	nodes 			map[string]*Node
	nodes 			map[int]*Node
	adjacentList 	map[string][]Edge
	distances		map[string]int
}

const MAXINT = 2147483647 // 0x7fffffff

func createNode(name string) *Node {
	return &Node{name, 0, MAXINT}
}

func (g *Graph) addNode(node *Node) {
	//g.nodes = append(g.nodes, node) 
	node.id = g.cnt
	g.cnt++
//	g.nodes[node.name] = node
	g.nodes[node.id] = node
	g.adjacentList[node.name] = []Edge{}
	g.distances[node.name] = MAXINT
}

func (g *Graph) addEdge(node1, node2 *Node, weight int) {
	g.adjacentList[node1.name] = append(g.adjacentList[node1.name], Edge{node2, weight})
	g.adjacentList[node2.name] = append(g.adjacentList[node2.name], Edge{node1, weight})
}

func (g *Graph) setPath(startingNode, finishingNode *Node) {
	g.startNode = startingNode
	g.endNode = finishingNode
	startingNode.cumul = 0
}

var fullstack, starbucks, insomnia, cafe, dig, dubliner *Node

func (g *Graph) Build() {

	g.addNode(fullstack)
	g.addNode(starbucks)
	g.addNode(insomnia)
	g.addNode(cafe)
	g.addNode(dig)
	g.addNode(dubliner)
	
	g.addEdge(fullstack, starbucks, 6);
	g.addEdge(fullstack, dig, 7);
	g.addEdge(fullstack, dubliner, 2);

	g.addEdge(dig, dubliner, 4);
	g.addEdge(dig, cafe, 9);

	g.addEdge(cafe, insomnia, 5);

	g.addEdge(insomnia, starbucks, 6);
	g.addEdge(insomnia, dubliner, 7);

	g.addEdge(starbucks, dubliner, 3);
}

