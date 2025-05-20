package egn

import "fmt"

type Show interface {
	MarginalEvents()
	// ConditionalEvents()
	// JointEvents()
}

func (n *Node) MarginalEvents() {
	if len(n.Marginal.Space) > 0 {
		for event, probability := range n.Marginal.Space {
			fmt.Println(event, ": ", probability)
		}
	}
}
