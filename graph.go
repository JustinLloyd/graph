package graph

type Node struct {
	Value interface{}
}

type Edge struct {
	From, To *Node
	Weight   float64
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func (g *Graph) AddNode(value interface{}) *Node {
	node := &Node{Value: value}
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *Graph) AddEdge(from, to *Node, weight float64) *Edge {
	edge := &Edge{From: from, To: to, Weight: weight}
	g.Edges = append(g.Edges, edge)
	return edge
}

func (g *Graph) OutgoingEdges(node *Node) []*Edge {
	var outgoingEdges []*Edge
	for _, edge := range g.Edges {
		if edge.From == node {
			outgoingEdges = append(outgoingEdges, edge)
		}
	}
	return outgoingEdges
}

func (g *Graph) IncomingEdges(node *Node) []*Edge {
	var incomingEdges []*Edge
	for _, edge := range g.Edges {
		if edge.To == node {
			incomingEdges = append(incomingEdges, edge)
		}
	}
	return incomingEdges
}
