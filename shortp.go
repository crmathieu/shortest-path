package main
import (
	"fmt"
)

type Edge struct {
    toNode  string
    weight  int
}
type NodeMapping map[string][]Edge

//type Node struct {
//    name string
//    neighbors []Edge
//}

type Graph struct {
    nodes []string
    //adjacentList []Edge
	nodeMap NodeMapping 
}


func (g *Graph) addNode(node string) {
	g.nodes = append(g.nodes, node) 
	g.nodeMap[node] = []Edge{}
}

func (g *Graph) addEdge(node1, node2 string, weight int) {
	g.nodeMap[node1] = append(g.nodeMap[node1], Edge{node2, weight})
	g.nodeMap[node2] = append(g.nodeMap[node2], Edge{node1, weight})
}

func NewGraph() *Graph {
	return &Graph{[]string{}, NodeMapping{}}
}

func (g *Graph) Build() {

	g.addNode("Fullstack")
	g.addNode("Starbucks")
	g.addNode("Insomnia Cookies")
	g.addNode("Cafe Grumpy")
	g.addNode("Dig Inn")
	g.addNode("Dubliner")
	
	g.addEdge("Fullstack", "Starbucks", 6);
	g.addEdge("Fullstack", "Dig Inn", 7);
	g.addEdge("Fullstack", "Dubliner", 2);

	g.addEdge("Dig Inn", "Dubliner", 4);
	g.addEdge("Dig Inn", "Cafe Grumpy", 9);

	g.addEdge("Cafe Grumpy", "Insomnia Cookies", 5);

	g.addEdge("Insomnia Cookies", "Starbucks", 6);
	g.addEdge("Insomnia Cookies", "Dubliner", 7);

	g.addEdge("Starbucks", "Dubliner", 3);
}

func main() {
	g := NewGraph()
	g.Build()
}