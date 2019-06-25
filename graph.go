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
//	cnt                int
	nodes              map[int]*Node
	neighborsList      map[int][]Edge
}

func NewGraph() *Graph {
	neighbors := make(map[int][]Edge)
	nodes := make(map[int]*Node)

//	return &Graph{0, nodes, neighbors}
	return &Graph{nodes, neighbors}
}

/*func createNode(name string) *Node {
	return &Node{name, 0}
}*/

//func (g *Graph) addNode(node *Node) {
func (g *Graph) addNode(nodename string, nodeid int) { //*Node {

	node := Node{}
	node.id = nodeid
	node.name = nodename
	g.nodes[node.id] = &node
	g.neighborsList[node.id] = []Edge{}
	//return &node
}

func (g *Graph) addEdge(node1, node2 *Node, weight int) {
	g.neighborsList[node1.id] = append(g.neighborsList[node1.id], Edge{node2, weight})
	g.neighborsList[node2.id] = append(g.neighborsList[node2.id], Edge{node1, weight})
}

var fullstack, starbucks, insomnia, cafe, dig, dubliner *Node
var jGraph *jsonGraph 

func (g *Graph) buildGraph(filename string) {

	jGraph = ReadJsonGraph(filename)
	for _, v := range(jGraph.Nodes) {
		g.addNode(v.Name, v.Id)
	}

	for _, v := range(jGraph.Nodes) {
		for _, z := range v.Neighbors {
			g.addEdge(g.nodes[v.Id], g.nodes[z.Id], z.Weight)
		}
	}
}

/*
func (g *Graph) BuildGraph2() {

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
*/