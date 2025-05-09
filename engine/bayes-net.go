package engine

import "fmt"

type Node struct {
	name     string
	marg     *ProbabilitySpace
	cond     *ProbabilitySpace
	joint    *ProbabilitySpace
	parents  []*Node
	children []*Node
	cpt      map[string]float64
}

type BayesNet struct {
	Nodes map[string]*Node
}

func NewNode(name string) *Node {
	return &Node{
		name:     name,
		marg:     NewProbabilitySpace(),
		cond:     NewProbabilitySpace(),
		joint:    NewProbabilitySpace(),
		parents:  []*Node{},
		children: []*Node{},
		cpt:      make(map[string]float64),
	}
}

func (n *Node) SetMarginal(event string, prob float64) {
	if len(n.parents) > 0 {
		fmt.Println("Error: you can't specify marginal probability since the node", n.name, "these parents")
		for _, parent := range n.parents {
			fmt.Print(parent.name, " ")
		}
		fmt.Printf("\n")

		return
	}

	n.marg.AddPair(event, prob)
}
