package egn

import "fmt"

type Show interface {
	MarginalEvents()
	ConditionalEvents()
}

func (n *Node) MarginalEvents() {
	if len(n.Marginal.Space) > 0 {
		for event, probability := range n.Marginal.Space {
			fmt.Println(event, ":", probability)
		}
		return

	}
	fmt.Println("no marginal event exists")
}

func (n *Node) ConditionalEvents() {
	if len(n.Conditional) > 0 {
		for dependencyCombination, space := range n.Conditional {
			for event, probability := range space.Space {
				fmt.Println(dependencyCombination, ":", event, ":", probability)
			}
		}
	}
}
