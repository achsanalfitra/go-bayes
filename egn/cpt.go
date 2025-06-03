package egn

import "fmt"

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
	combinationsIndex := make([]int, 0, len(n.Parents))

	for _, parent := range n.Parents {
		combinationsIndex = append(combinationsIndex, len(parent.States.StrInt))
	}

	return combinationsIndex
}

func (n *Node) CreateCombinations(input int) ([]int, error) {
	// return error if out of range
	if input > n.TotalStates()-1 {
		return nil, fmt.Errorf("input out of range, the maximum is %d", n.TotalStates()-1)
	}

	combinationIndex := n.CombinationsIndex()
	output := make([]int, 0, len(n.Parents))

	// iterate from 0 to totalcombinations length - 1
	for i := 0; i < len(combinationIndex); i++ {

		// append output with modulo
		output = append(output, input%combinationIndex[i])
		input /= combinationIndex[i]
	}

	return output, nil
}

func (n *Node) CombinationToIndex(combination []int) (int, error) {
	combinationIndex := n.CombinationsIndex()

	if len(combination) != len(combinationIndex) {
		return 0, fmt.Errorf("expected %d, got %d", len(combinationIndex), len(combination))
	}

	index := 0
	multiplier := 1

	for i := 0; i < len(combination); i++ {
		if combination[i] >= combinationIndex[i] {
			return 0, fmt.Errorf("combination[%d] = %d exceeds max state %d", i, combination[i], combinationIndex[i]-1)
		}
		index += combination[i] * multiplier
		multiplier *= combinationIndex[i]
	}

	return index, nil
}
