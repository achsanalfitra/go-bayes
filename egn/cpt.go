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

// func (n *Node) GenerateCPT() {
// 	parentCombinationMap := make(map[string]string)
// 	ownState := n.States
// 	for
// }
