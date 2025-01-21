package template

import (
	"fmt"
	"sync"
)

type Node struct {
	val int
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

func (g *Graph) AddEdge(u, v *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*u] = append(g.edges[*u], v)
	g.edges[*v] = append(g.edges[*v], u)
}

func (g *Graph) String() string {
	g.lock.Lock()
	defer g.lock.Unlock()

	str := ""
	for _, iNode := range g.nodes {
		str += iNode.String() + " -> "
		nexts, _ := g.edges[*iNode]
		for _, next := range nexts {
			str += next.String() + " "
		}
		str += "\n"
	}
	return str
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.val)
}