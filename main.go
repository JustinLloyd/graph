package graph

import "fmt"

func main() {
	g := &Graph{}
	node1 := g.AddNode("A")
	node2 := g.AddNode("B")
	g.AddEdge(node1, node2, 10)

	// Print graph
	for _, edge := range g.Edges {
		fmt.Printf("%v -> %v (Weight: %v)\n", edge.From.Value, edge.To.Value, edge.Weight)
	}
}
