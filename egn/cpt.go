package egn

type CPT struct {
	Missing  map[string]struct{}
	Existing map[string]struct{}
}

func CreateCPT() *CPT {
	return &CPT{
		Missing:  make(map[string]struct{}),
		Existing: make(map[string]struct{}),
	}
}

func (n *Node) CombinationsIndex() []int {
	// assign parents length as index
	parentsLength := len(n.Parents)
	combinationsIndex := make([]int, 0, parentsLength)

	for _, parent := range n.Parents {
		combinationsIndex = append(combinationsIndex, len(parent.States))
	}

	return combinationsIndex
}
