# Graph Package

The Graph package offers a small flexible toolkit for working with graphs in Go. Whether you're dealing with directed or
undirected, weighted or unweighted graphs, this package provides a streamlined interface to create, manipulate, and analyze
graph structures. It forms part of a larger framework for working with my FFMPEG library wrapper which is also written in Go and
my gstreamer library wrapper, again, also written in Go. 

Designed with simplicity in mind, the Graph package enables developers to perform common graph operations 
such as depth-first search (DFS), breadth-first search (BFS), topological sorting, cycle detection, and pathfinding algorithms
like A*.

The Graph package is built to be simple and efficient, with no extraneous functions other than those needed for the purpose
of manipulating ffmpeg and gstreamer graphs. It has a modular and extensible design to ensure that developers can easily
integrate it into their projects and customize it to fit specific needs.

Explore the sections below to dive into the features, functions, and usage examples of the Graph package.


## Types

### Graph
```go
type Node struct {
	Name   string
	Object interface{}
}
```

### Node

Represents a node in the graph. It has a name and can carry an associated object.

```go
type Node struct {
	Name   string
	Object interface{}
}
```

### Edge

Represents an edge between two nodes with an optional weight.

```go
type Edge struct {
	From     *Node
	To       *Node
	Weight   float64
	Directed bool
}
```

### EdgeType

An enum representing different types of edges: Tree Edge, Back Edge, Forward Edge, and Cross Edge.

```go
type EdgeType int
```

## Functions

### AddNode

Adds a new node to the graph.

```go
func (g *Graph) AddNode(name string, object interface{}) *Node {}
```

### AddEdge

Adds a new edge to the graph.

```go
func (g *Graph) AddEdge(from, to *Node, weight float64, directed bool) *Edge {}
```

### OutgoingEdges

Get all of the outgoing edges for a specific node in the graph.

```go
func (g *Graph) OutgoingEdges(node *Node) []*Edge {}
```

### IncomingEdges

Get all of the incoming edges for a specific node in the graph.

```go
func (g *Graph) IncomingEdges(node *Node) []*Edge {}
```

### Neighbours

Get all of the neighbours of a specific node.

```go
func (g *Graph) Neighbors(node *Node) []*Node {}
```

### DFS (depth-first search)
Performs a depth-first search of a graph.

```go
func (g *Graph) DFS(start *Node, visited map[*Node]bool, process func(*Node)) {}
```

### DetectCycles

Determines if any cycles exist in the graph, returns a list of nodes that have them.

```go
func (g *Graph) DetectCycles(start *Node) [][]*Node {}
```

### TopologicalSort
Topologically sorts the graph, used for scheduling, windowing overlaps and dependency determination in gstreamer.

```go
func (g *Graph) TopologicalSort() []*Node {}
```

### BFS

Performs a simple breadth-first search for reachability testing.

```go
func (g *Graph) BFS(start *Node, process func(*Node)) {}
```

### ConnectedComponents

Gets all of the connected components. Come on, read a book on graph theory...

```go
func (g *Graph) ConnectedComponents() [][]*Node {}
```

### ClassifyEdges

Determines what an edge is, leaf, branch, etc. Used for debugging and diagnosing gstreamer graphs.

```go
func (g *Graph) ClassifyEdges() map[*Edge]EdgeType {}
```

### IsDAG

Determines if the graph is directed or not. Useful for debugging gstreamer graphs.

```go
func (g *Graph) IsDAG() bool {}
```

### IsWeighted

Used to determine if the graph contains weights, which determines if an ffmpeg or gstreamer operation should be done on the GPU or the CPU, and single-threaded or multi-threaded.

```go
func (g *Graph) IsWeighted() bool {}
```

### AStar

Oddly enough, A* is useful for figuring out where to send a gstreamer pipeline operation. "How expensive is this operation? Should I rewrite the pipeline?"

```go
func (g *Graph) AStar(start *Node, goal *Node) []*Node {}
```

### TopologicalSort

Returns a topological sort of the graph.

```go
func (g *Graph) TopologicalSort() []*Node {}
```

### AStar

Implements the A* algorithm to find the shortest path between two nodes.

```go
func (g *Graph) AStar(start *Node, goal *Node) []*Node {}
```

## Usage

You can create a graph and add nodes and edges to it, and then perform various graph operations such as depth-first search, breadth-first search, cycle detection, and more.

## Installation

Include the following import statement in your Go file:

```go
import "justinlloyd.com/graph"
```

## Support

Absolutely none provided unless you are paying me $200/hr with a six-month contract.

## License

    ** DO WHATEVER PUBLIC LICENSE **

    ** TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION**

    0. You can do whatever you want to with the work.
    1. You cannot stop anybody from doing whatever they want to with the work.
    2. You cannot revoke anybody elses DO WHATEVER PUBLIC LICENSE in the work.

    This program is free software. It comes without any warranty, to
    the extent permitted by applicable law. You can redistribute it
    and/or modify it under the terms of the DO WHATEVER PUBLIC LICENSE

## Contributing

Don't. I don't want to know about it. I don't care.

Software originally created by Justin Lloyd @ http://justin-lloyd.com/


