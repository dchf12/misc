package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	name string
	cost float64
}

type NodeHeap []*Node

var _ heap.Interface = (*NodeHeap)(nil)

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x any) {
	node := x.(*Node)
	*h = append(*h, node)
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[:n-1]
	return node
}

type Edge struct {
	from *Node
	to   *Node
	cost float64
}

type Graph struct {
	nodes map[string]*Node
	edges map[string][]*Edge
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
		edges: make(map[string][]*Edge),
	}
}

func (g *Graph) AddNode(node *Node) {
	g.nodes[node.name] = node
}

func (g *Graph) AddEdge(from, to *Node, cost float64) {
	edge := &Edge{
		from: from,
		to:   to,
		cost: cost,
	}

	g.edges[from.name] = append(g.edges[from.name], edge)
}

// Dijkstra returns the shortest path from start to all other nodes in the graph.
func (g *Graph) Dijkstra(start *Node) map[string]float64 {
	distances := make(map[string]float64)
	for node := range g.nodes {
		distances[node] = math.Inf(1)
	}
	distances[start.name] = 0

	pq := make(NodeHeap, 1)
	pq[0] = start
	start.cost = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		minNode := heap.Pop(&pq).(*Node)
		fmt.Println("pop", minNode.name, minNode.cost)
		for _, edge := range g.edges[minNode.name] {
			alt := minNode.cost + edge.cost
			fmt.Println("search", edge.to.name, alt)
			if alt < distances[edge.to.name] {
				distances[edge.to.name] = alt
				edge.to.cost = alt
				heap.Push(&pq, edge.to)
				fmt.Println("push", edge.to.name, edge.to.cost)
			}
		}
	}
	return distances
}

func main() {
	graph := NewGraph()

	nodeA := &Node{name: "A"}
	nodeB := &Node{name: "B"}
	nodeC := &Node{name: "C"}
	nodeD := &Node{name: "D"}
	nodeE := &Node{name: "E"}

	graph.AddNode(nodeA)
	graph.AddNode(nodeB)
	graph.AddNode(nodeC)
	graph.AddNode(nodeD)
	graph.AddNode(nodeE)

	graph.AddEdge(nodeA, nodeB, 6)
	graph.AddEdge(nodeA, nodeC, 1)
	graph.AddEdge(nodeB, nodeD, 5)
	graph.AddEdge(nodeC, nodeD, 2)
	graph.AddEdge(nodeC, nodeE, 4)

	distances := graph.Dijkstra(nodeA)
	for node, cost := range distances {
		fmt.Println(node, cost)
	}
}
