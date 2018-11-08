package graph

import "fmt"

// Item is a generic type for graph
type Item interface{}

// Node is a point on the graph
type Node struct {
	value Item
}

// MakeNode creates a node out of data input
func MakeNode(item Item) *Node {
	return &Node{item}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// Graph is an adjacency list undirected implementation of the graph object
type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

// AddNode places a new node into the graph
func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

// AddEdge connects two node edges
func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}

	g.edges[*n1] = append(g.edges[*n1], n2)

	g.edges[*n2] = append(g.edges[*n2], n1)

}

func (g *Graph) String() string {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}

	return s
}
