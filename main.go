package main

import "github.com/Sagardhandhalya/palgo/graph"

func main() {
	// 1. graph test
	g := graph.Graph{}
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(3, 2)
	g.PrintGraph()
}
