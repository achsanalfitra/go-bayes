package engine

import (
	"fmt"
	"slices"
)

type Node struct {
	name     string
	states   []string
	marg     *ProbabilitySpace
	cond     *ProbabilitySpace
	joint    *ProbabilitySpace
	parents  []*Node
	children []*Node
	cpt      map[string]float64
	margSet  bool
	condSet  bool
	jointSet bool
}

type BayesNet struct {
	Nodes map[string]*Node
}

func NewNode(name string) *Node {
	return &Node{
		name:     name,
		states:   []string{},
		marg:     NewProbabilitySpace(),
		cond:     NewProbabilitySpace(),
		joint:    NewProbabilitySpace(),
		parents:  []*Node{},
		children: []*Node{},
		cpt:      make(map[string]float64),
		margSet:  false,
		condSet:  false,
		jointSet: false,
	}
}

func (n *Node) SetMarg(event string, prob float64) {
	if len(n.parents) > 0 {
		fmt.Println("Error: you can't specify marginal probability since the node", n.name, "these parents")
		for _, parent := range n.parents {
			fmt.Print(parent.name, " ")
		}
		fmt.Printf("\n")

		return
	}

	n.marg.AddPair(event, prob)
	n.UpdateState(event)
	n.margSet = true
}

func (n *Node) SetCond(event string, prob float64) {
	if len(n.parents) == 0 {
		fmt.Println("Error: you can't specify conditional probability since the node", n.name, "has no parents")

		return
	}

	if n.margSet && n.jointSet {
		fmt.Println("Error: you can't specify conditional probability since the node", n.name, "already has marginal and joint probability specified")

		return
	}

	n.cond.AddPair(event, prob)
	n.UpdateState(event)
	n.condSet = true
}

func (n *Node) UpdateState(event string) {
	if !slices.Contains(n.states, event) {
		n.states = append(n.states, event)
	}
}
