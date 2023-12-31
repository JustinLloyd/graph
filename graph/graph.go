package graph

import "math"

type EdgeType int

const (
	TreeEdge EdgeType = iota
	BackEdge
	ForwardEdge
	CrossEdge
)

type Node struct {
	Name   string
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

func (g *Graph) AddNode(name string, object interface{}) *Node {
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
	g.dfsFindCycles(start, visited, path, &cycles)
	return cycles
}

func (g *Graph) dfsFindCycles(current *Node, visited map[*Node]bool, path []*Node, cycles *[][]*Node) {
	visited[current] = true
	path = append(path, current)

	for _, neighbor := range g.Neighbors(current) {
		if !visited[neighbor] {
			g.dfsFindCycles(neighbor, visited, path, cycles)
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

func (g *Graph) TopologicalSort() []*Node {
	stack := []*Node{}
	visited := make(map[*Node]bool)

	// Visit all the vertices one by one
	for _, node := range g.Nodes {
		if !visited[node] {
			g.dfsTopologicalSort(node, visited, &stack)
		}
	}

	return reverse(stack)
}

func (g *Graph) dfsTopologicalSort(v *Node, visited map[*Node]bool, stack *[]*Node) {
	// Mark the current node as visited
	visited[v] = true

	// Recur for all the vertices adjacent to this vertex
	for _, neighbor := range g.Neighbors(v) {
		if !visited[neighbor] {
			g.dfsTopologicalSort(neighbor, visited, stack)
		}
	}

	// Push current vertex to stack which stores result
	*stack = append(*stack, v)
}

func reverse(nodes []*Node) []*Node {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
	return nodes
}

func (g *Graph) BFS(start *Node, process func(*Node)) {
	visited := make(map[*Node]bool)
	queue := []*Node{start}

	for len(queue) > 0 {
		// Dequeue a vertex from queue and process it
		node := queue[0]
		queue = queue[1:]
		if visited[node] {
			continue
		}
		visited[node] = true
		process(node)

		// Get all adjacent vertices of the dequeued vertex
		// If an adjacent vertex is not visited, then mark it visited and enqueue it
		for _, neighbor := range g.Neighbors(node) {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func (g *Graph) ConnectedComponents() [][]*Node {
	visited := make(map[*Node]bool)
	var components [][]*Node

	for _, node := range g.Nodes {
		if !visited[node] {
			var component []*Node
			g.DFS(node, visited, func(n *Node) {
				component = append(component, n)
			})
			components = append(components, component)
		}
	}

	return components
}

func (et EdgeType) String() string {
	switch et {
	case TreeEdge:
		return "Tree Edge"
	case BackEdge:
		return "Back Edge"
	case ForwardEdge:
		return "Forward Edge"
	case CrossEdge:
		return "Cross Edge"
	}

	return "Unknown Edge Type"
}

func (g *Graph) FindEdge(from, to *Node) *Edge {
	for _, edge := range g.Edges {
		if edge.Directed {
			if edge.From == from && edge.To == to {
				return edge
			}
		} else {
			if (edge.From == from && edge.To == to) || (edge.From == to && edge.To == from) {
				return edge
			}
		}
	}

	return nil // Return nil if no such edge is found
}

func (g *Graph) ClassifyEdges() map[*Edge]EdgeType {
	classification := make(map[*Edge]EdgeType)
	visited := make(map[*Node]bool)
	timestamps := make(map[*Node][]int)
	time := 0

	var dfs func(*Node)
	dfs = func(node *Node) {
		visited[node] = true
		timestamps[node] = []int{time, 0}
		time++

		for _, neighbor := range g.Neighbors(node) {
			edge := g.FindEdge(node, neighbor)
			if !visited[neighbor] {
				classification[edge] = TreeEdge
				dfs(neighbor)
			} else if timestamps[neighbor][0] > timestamps[node][0] {
				classification[edge] = ForwardEdge
			} else if timestamps[neighbor][1] == 0 {
				classification[edge] = BackEdge
			} else {
				classification[edge] = CrossEdge
			}
		}

		timestamps[node][1] = time
		time++
	}

	for _, node := range g.Nodes {
		if !visited[node] {
			dfs(node)
		}
	}

	return classification
}

func (g *Graph) IsDAG() bool {
	visited := make(map[*Node]bool)
	onStack := make(map[*Node]bool) // Keep track of nodes currently on the DFS stack

	var dfs func(node *Node) bool
	dfs = func(node *Node) bool {
		visited[node] = true
		onStack[node] = true

		for _, neighbor := range g.Neighbors(node) {
			// If the neighbor is on the current DFS stack, a cycle is detected
			if onStack[neighbor] {
				return false
			}

			if !visited[neighbor] && !dfs(neighbor) {
				return false
			}
		}

		onStack[node] = false // Remove node from DFS stack
		return true
	}

	for _, node := range g.Nodes {
		if !visited[node] && !dfs(node) {
			return false
		}
	}

	return true
}

func (g *Graph) IsWeighted() bool {
	for _, edge := range g.Edges {
		// Check if the weight of the edge is different from the default (e.g., 0)
		// If so, it indicates that the graph is weighted
		if edge.Weight != 0 {
			return true
		}
	}
	return false
}

func (g *Graph) AStar(start *Node, goal *Node) []*Node {
	openSet := []*Node{start}
	cameFrom := make(map[*Node]*Node)
	gScore := make(map[*Node]float64)
	gScore[start] = 0
	fScore := make(map[*Node]float64)
	fScore[start] = gScore[start] // Assuming heuristic cost from start to goal is 0

	for len(openSet) > 0 {
		current := openSet[0]
		minFScore := fScore[current]

		for _, node := range openSet[1:] {
			score := fScore[node]
			if score < minFScore {
				current = node
				minFScore = score
			}
		}

		if current == goal {
			return reconstructPath(cameFrom, current)
		}

		openSet = removeFromSet(openSet, current)

		neighbors := g.Neighbors(current)
		for _, neighbor := range neighbors {
			edge := g.FindEdge(current, neighbor)

			// Check for directed edge when graph is not a DAG
			if edge != nil && (edge.Directed == false || edge.From == current) {
				tentativeGScore := gScore[current] + edge.Weight
				if tentativeGScore < gScore[neighbor] {
					cameFrom[neighbor] = current
					gScore[neighbor] = tentativeGScore
					fScore[neighbor] = tentativeGScore
					if !inSet(openSet, neighbor) {
						openSet = append(openSet, neighbor)
					}
				}
			}
		}
	}

	return nil
}

func heuristic(node, goal *Node, edges []*Edge) float64 {
	minWeight := math.Inf(1)
	for _, edge := range edges {
		if (edge.From == node && edge.To == goal) || (edge.To == node && edge.From == goal) {
			if edge.Weight < minWeight {
				minWeight = edge.Weight
			}
		}
	}

	return minWeight
}

func distance(from, to *Node, edges []*Edge) float64 {
	for _, edge := range edges {
		if (edge.From == from && edge.To == to) || (edge.To == from && edge.From == to) {
			return edge.Weight
		}
	}
	return math.Inf(1)
}

func reconstructPath(cameFrom map[*Node]*Node, current *Node) []*Node {
	path := []*Node{current}
	for {
		previous, exists := cameFrom[current]
		if !exists {
			break
		}
		path = append([]*Node{previous}, path...)
		current = previous
	}
	return path
}

func removeFromSet(set []*Node, node *Node) []*Node {
	for i, n := range set {
		if n == node {
			return append(set[:i], set[i+1:]...)
		}
	}
	return set
}

func inSet(set []*Node, node *Node) bool {
	for _, n := range set {
		if n == node {
			return true
		}
	}
	return false
}
