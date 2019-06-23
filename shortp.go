package main
import (
	"fmt"
)

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
	nodes := make(map[string]*Node)
//	return &Graph{&Node{}, &Node{}, []*Node{}, adj}
	return &Graph{&Node{}, &Node{}, 0, nodes, adj, dist}
}

type Graph struct {
	startNode, endNode  		*Node
	cnt 			int
	nodes 			map[string]*Node
//	nodes			[]*Node
	adjacentList 	map[string][]Edge
	distances		map[string]int
}

func (g *Graph) findShortest(start, end *Node) {

	keepGoing := true
	for n := start; keepGoing; {
		edges := g.adjacentList[n.name]
		//least := MAXINT
		for _, edge := range(edges) {
			/*if dist, ok := g.distances[edge.toNode.name]; ok {
				if dist > edge.weight + n.cumul {
					g.distances[edge.toNode.name] = edge.weight + n.cumul
				}
			}*/
			if edge.toNode.cumul == MAXINT {
				// first time we reach this node
				edge.toNode.cumul = edge.weight + n.cumul
			}
			
			if edge.toNode.cumul > edge.weight + n.cumul {
				edge.toNode.cumul = edge.weight + n.cumul

			}
			if edge.toNode.name == end.name {

			}
			
		}
	}
}


const MAXINT = 2147483647 // 0x7fffffff

func createNode(name string) *Node {
	return &Node{name, 0, MAXINT}
}

func (g *Graph) addNode(node *Node) {
	//g.nodes = append(g.nodes, node) 
	node.id = g.cnt
	g.cnt++
	g.nodes[node.name] = node
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

var fullstack, starbucks, insomnia, cafe, dig, dubliner *Node

func main() {
	fullstack 	= createNode("Fullstack")
	starbucks 	= createNode("Starbucks")
	insomnia 	= createNode("Insomnia Cookies")
	cafe 		= createNode("Cafe Grumpy")
	dig 		= createNode("Dig Inn")
	dubliner 	= createNode("Dubliner")

	g := NewGraph()
	g.Build()
	g.setPath(fullstack, cafe)
	fmt.Println(g.nodes)
/*
	s := NewStack()
	s.Push(&Item{1})
	s.Push(&Item{2})
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(&Item{3})
	fmt.Println(s.Pop())*/
}
