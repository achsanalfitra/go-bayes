package egn

type NodeValidator interface {
	TotalStates() int
	ConditionalValid() bool
	MarginalValid() bool
}

func (n *Node) TotalStates() int {
	// calculate own states
	ownState := len(n.States)

	// set default parent states
	parentStates := 1

	for _, parent := range n.Parents {
		parentStates *= len(parent.States)
	}

	// correct cartesian product formula own state * product of parent states
	ownState *= parentStates

	return ownState
}
