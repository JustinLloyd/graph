package graph

type Node struct {
	Name   *string
	Object interface{}
}

type Edge struct {
	From     *Node
	To       *Node
	Weight   float64
	Directed bool
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func (g *Graph) AddNode(name *string, object interface{}) *Node {
	node := &Node{Name: name, Object: object}
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *Graph) AddEdge(from, to *Node, weight float64, directed bool) *Edge {
	edge := &Edge{From: from, To: to, Weight: weight, Directed: directed}
	g.Edges = append(g.Edges, edge)
	if !directed {
		reverseEdge := &Edge{From: to, To: from, Weight: weight, Directed: false}
		g.Edges = append(g.Edges, reverseEdge)
	}

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

func (g *Graph) Neighbors(node *Node) []*Node {
	neighborsMap := make(map[*Node]bool)
	for _, edge := range g.Edges {
		if edge.From == node {
			neighborsMap[edge.To] = true
		}
		if edge.To == node {
			neighborsMap[edge.From] = true
		}
	}

	var neighbors []*Node
	for neighbor := range neighborsMap {
		neighbors = append(neighbors, neighbor)
	}

	return neighbors
}

func (g *Graph) DFS(start *Node, visited map[*Node]bool, process func(*Node)) {
	// Mark the current node as visited
	visited[start] = true

	// Process the current node (e.g., print its value or name)
	process(start)

	// Recur for all the neighbors of this node
	for _, neighbor := range g.Neighbors(start) {
		if !visited[neighbor] {
			g.DFS(neighbor, visited, process)
		}
	}
}

func (g *Graph) DetectCycles(start *Node) [][]*Node {
	visited := make(map[*Node]bool)
	path := []*Node{}
	cycles := [][]*Node{}
	g.findCyclesDFS(start, visited, path, &cycles)
	return cycles
}

func (g *Graph) findCyclesDFS(current *Node, visited map[*Node]bool, path []*Node, cycles *[][]*Node) {
	visited[current] = true
	path = append(path, current)

	for _, neighbor := range g.Neighbors(current) {
		if !visited[neighbor] {
			g.findCyclesDFS(neighbor, visited, path, cycles)
		} else if isInPath(neighbor, path) {
			cycle := append([]*Node{}, path...) // Copy the current path
			cycle = append(cycle, neighbor)     // Add the current neighbor to close the cycle
			*cycles = append(*cycles, cycle)
		}
	}

	visited[current] = false // Backtrack from the current node
}

func isInPath(node *Node, path []*Node) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}
