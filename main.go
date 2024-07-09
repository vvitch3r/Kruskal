package main

import (
	"fmt"
	"sort"
)

// Edge represents an edge in the graph
type Edge struct {
	from, to, weight int
}

// Graph represents a weighted undirected graph using an edge list
type Graph struct {
	vertices int
	edges    []Edge
}

// NewGraph initializes a new graph with a given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
	}
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.edges = append(g.edges, Edge{from, to, weight})
}

// DisjointSet represents a disjoint-set data structure (Union-Find)
type DisjointSet struct {
	parent, rank []int
}

// NewDisjointSet initializes a new disjoint-set with a given number of elements
func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DisjointSet{parent, rank}
}

// Find finds the representative of the set containing element x
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) // Path compression
	}
	return ds.parent[x]
}

// Union unites the sets containing elements x and y
func (ds *DisjointSet) Union(x, y int) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)
	if rootX != rootY {
		// Union by rank
		if ds.rank[rootX] > ds.rank[rootY] {
			ds.parent[rootY] = rootX
		} else if ds.rank[rootX] < ds.rank[rootY] {
			ds.parent[rootX] = rootY
		} else {
			ds.parent[rootY] = rootX
			ds.rank[rootX]++
		}
	}
}

// KruskalMST finds the Minimum Spanning Tree using Kruskal's algorithm
func (g *Graph) KruskalMST() int {
	// Sort the edges by weight
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	// Create a disjoint-set for tracking connected components
	ds := NewDisjointSet(g.vertices)

	// Total weight of the MST
	totalWeight := 0

	// Iterate through sorted edges and add them to the MST
	for _, edge := range g.edges {
		if ds.Find(edge.from) != ds.Find(edge.to) {
			ds.Union(edge.from, edge.to)
			totalWeight += edge.weight
		}
	}

	return totalWeight
}

func main() {
	// Create a new graph with 5 vertices
	graph := NewGraph(5)

	// Add edges to the graph
	graph.AddEdge(0, 1, 2)
	graph.AddEdge(0, 3, 6)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(1, 3, 8)
	graph.AddEdge(1, 4, 5)
	graph.AddEdge(2, 4, 7)
	graph.AddEdge(3, 4, 9)

	// Find the MST using Kruskal's algorithm
	totalWeight := graph.KruskalMST()

	// Print the total weight of the MST
	fmt.Printf("Total weight of the Minimum Spanning Tree: %d\n", totalWeight)
}
