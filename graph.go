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
	nodes              map[int]*Node
	neighborsList      map[int][]Edge
}

func NewGraph() *Graph {
	neighbors := make(map[int][]Edge)
	nodes := make(map[int]*Node)
	return &Graph{nodes, neighbors}
}

//func (g *Graph) addNode(node *Node) {
func (g *Graph) addNode(nodename string, nodeid int) { //*Node {

	node := Node{}
	node.id = nodeid
	node.name = nodename
	g.nodes[node.id] = &node
	g.neighborsList[node.id] = []Edge{}
}

func (g *Graph) addEdge(node1, node2 *Node, weight int) {
	g.neighborsList[node1.id] = append(g.neighborsList[node1.id], Edge{node2, weight})
	g.neighborsList[node2.id] = append(g.neighborsList[node2.id], Edge{node1, weight})
}

//var fullstack, starbucks, insomnia, cafe, dig, dubliner *Node
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
