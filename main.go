package main

import (
	"fmt"

	"github.com/Sagardhandhalya/palgo/graph"
	"github.com/Sagardhandhalya/palgo/tree"
)

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

	// 2. tree
	bt := tree.Node{Val: 10, Left: nil, Right: nil}
	fmt.Println(bt)
	tree.Insert(&bt, 15)
	tree.Insert(&bt, 12)
	tree.Insert(&bt, 18)
	tree.Insert(&bt, 5)
	tree.Insert(&bt, 3)
	tree.Insert(&bt, 6)
	tree.InOrder(&bt)
}
