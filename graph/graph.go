package graph

import "fmt"

type Graph struct {
	Vertexts []*Node
}

type Node struct {
	Key  int
	Conn []*Node
}

func (g *Graph) AddVertex(k int) {
	g.Vertexts = append(g.Vertexts, &Node{k, nil})
}

func keyToPonter(key int, g *Graph) *Node {
	for _, n := range g.Vertexts {
		if n.Key == key {
			return n
		}
	}
	return nil
}

func (g *Graph) AddEdge(from, to int) {
	pFrom := keyToPonter(from, g)
	pTo := keyToPonter(to, g)

	if pFrom == nil || pTo == nil {
		err := fmt.Errorf("invalid Edge (%d - %d )", from, to)
		fmt.Println(err.Error())
	} else {
		pFrom.Conn = append(pFrom.Conn, pTo)
	}
}

func (g *Graph) PrintGraph() {
	for _, n := range g.Vertexts {
		fmt.Printf("\n Vertext %d : ", n.Key)
		for _, c := range n.Conn {
			fmt.Printf("%d", c.Key)
		}
	}
}
