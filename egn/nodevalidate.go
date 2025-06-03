package egn

type NodeValidator interface {
	TotalStates() int
	ConditionalValid() bool
	MarginalValid() bool
}

func (n *Node) TotalStates() int {
	// calculate own states
	ownState := len(n.States.StrInt)

	// set default parent states
	parentStates := 1

	for _, parent := range n.Parents {
		parentStates *= len(parent.States.StrInt)
	}

	// correct cartesian product formula own state * product of parent states
	ownState *= parentStates

	return ownState
}

func (n *Node) ConditionalValid() bool {
	// find out own total states
	totalStates := n.TotalStates()
	ownConditionalStates := 0

	// sum all of own states
	for _, ps := range n.Conditional {
		ownConditionalStates += len(ps.Space)
	}

	// return size check
	return ownConditionalStates == totalStates
}

func (n *Node) MarginalValid() bool {
	// TODO: implement functions to validate marginal space
	return true
}
