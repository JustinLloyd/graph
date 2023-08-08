package main

import (
	"fmt"
	"justinlloyd.com/graph"
)

func main() {
	g := &graph.Graph{}
	node1 := g.AddNode("A", nil)
	node2 := g.AddNode("B", nil)
	g.AddEdge(node1, node2, 10, false)

	// Print graph
	for _, edge := range g.Edges {
		fmt.Printf("%v -> %v (Weight: %v)\n", edge.From.Object, edge.To.Object, edge.Weight)
	}
}
