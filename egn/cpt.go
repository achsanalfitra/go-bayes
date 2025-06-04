package egn

import "fmt"

type CPT struct {
	Known   map[int]map[int]struct{} // faster and lighter map storing using int since the parent and states are already mapped somewhere else
	Ignored map[int]map[int]struct{}
	Rebuild map[int]map[int]struct{}
	node    *Node
}

func CreateCPT(n *Node) *CPT {
	return &CPT{
		Known:   make(map[int]map[int]struct{}),
		Ignored: make(map[int]map[int]struct{}),
		Rebuild: make(map[int]map[int]struct{}),
		node:    n,
	}
}

func (cpt *CPT) AddKnown(parentID, parentStateID int) error {

	err := cpt.CheckState(parentID, parentStateID, cpt.node)
	if err != nil {
		return err
	}

	if _, ok := cpt.Known[parentID]; !ok {
		cpt.Known[parentID] = make(map[int]struct{})
	}

	cpt.Known[parentID][parentStateID] = struct{}{}

	return nil
}

func (cpt *CPT) AddIgnore(parentID, parentStateID int) error {

	err := cpt.CheckState(parentID, parentStateID, cpt.node)
	if err != nil {
		return err
	}

	if _, ok := cpt.Ignored[parentID]; !ok {
		cpt.Ignored[parentID] = make(map[int]struct{})
	}

	cpt.Ignored[parentID][parentStateID] = struct{}{}

	return nil
}

func (cpt *CPT) CheckState(parentID, parentStateID int, n *Node) error {
	parentKey, exist := n.ParentsMap.IntStr[parentID]

	if !exist {
		return fmt.Errorf("parentID %d doesn't exist", parentID)
	}

	if _, exist := n.Parents[parentKey].States.IntStr[parentStateID]; !exist {
		return fmt.Errorf("the state doesn't exist")
	}

	return nil
}

// func (n *Node) CombinationsIndex() []int {
// 	// assign parents length as index
// 	combinationsIndex := make([]int, 0, len(n.Parents))

// 	for _, parent := range n.Parents {
// 		combinationsIndex = append(combinationsIndex, len(parent.States.StrInt))
// 	}

// 	return combinationsIndex
// }

// func (n *Node) CreateCombinations(input int) ([]int, error) {
// 	// return error if out of range
// 	if input > n.TotalStates()-1 {
// 		return nil, fmt.Errorf("input out of range, the maximum is %d", n.TotalStates()-1)
// 	}

// 	combinationIndex := n.CombinationsIndex()
// 	output := make([]int, 0, len(n.Parents))

// 	// iterate from 0 to totalcombinations length - 1
// 	for i := 0; i < len(combinationIndex); i++ {

// 		// append output with modulo
// 		output = append(output, input%combinationIndex[i])
// 		input /= combinationIndex[i]
// 	}

// 	return output, nil
// }

// func (n *Node) CombinationToIndex(combination []int) (int, error) {
// 	combinationIndex := n.CombinationsIndex()

// 	if len(combination) != len(combinationIndex) {
// 		return 0, fmt.Errorf("expected %d, got %d", len(combinationIndex), len(combination))
// 	}

// 	index := 0
// 	multiplier := 1

// 	for i := 0; i < len(combination); i++ {
// 		if combination[i] >= combinationIndex[i] {
// 			return 0, fmt.Errorf("combination[%d] = %d exceeds max state %d", i, combination[i], combinationIndex[i]-1)
// 		}
// 		index += combination[i] * multiplier
// 		multiplier *= combinationIndex[i]
// 	}

// 	return index, nil
// }
