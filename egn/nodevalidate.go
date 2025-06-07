package egn

import "fmt"

type NodeValidator interface {
	TotalStates() int
	ConditionalValid() bool
	MarginalValidator
}

type MarginalValidator interface {
	MarginalCoverage(inferenceType, query string) error
	MarginalInternal() (bool, string)
}

func (n *Node) MarginalCoverage(inferenceType, query string) error {
	// base case for coverage check is it is query driven

	// create a multi use-cases
	switch inferenceType {
	case "joint":
		jointMap := MultiEventToMap(query)
		marginalState := jointMap[n.Name]
		if _, exists := n.Marginal.Space[SingleEventToString(n.Name, marginalState)]; !exists {
			return fmt.Errorf("the node %s does not have the inquired state %s", n.Name, marginalState)
		}
	case "conditional":
		_, givenEventsMap := ConditionalToMap(query)

		givenEventsSize := len(givenEventsMap)

		// handle invalid query
		if givenEventsSize > 1 {
			return fmt.Errorf("there are %d elements in the given events; use other validation methods to query", givenEventsSize)
		}

		marginalState := givenEventsMap[n.Name]
		if _, exists := n.Marginal.Space[SingleEventToString(n.Name, marginalState)]; !exists {
			return fmt.Errorf("the inquired state %s doesn't exist", marginalState)
		}

	default:
		return fmt.Errorf("%s is an invalid inference type, use inferenceType object to query", inferenceType)
	}

	return nil
}

func (n *Node) MarginalInternal() (bool, string) {
	// update internal validity\
	n.Marginal.UpdateValidity()

	// check space validity
	if !n.Marginal.CheckValidity() {
		return false, fmt.Sprintf("the probability is: %f; probability must be within 0 to 1 range ", n.Marginal.TotalProb())
	}

	return true, "marginal space valid"
}

// func (n *Node) TotalStates() int {
// 	// calculate own states
// 	ownState := len(n.States.StrInt)

// 	// set default parent states
// 	parentStates := 1

// 	for _, parent := range n.Parents {
// 		parentStates *= len(parent.States.StrInt)
// 	}

// 	// correct cartesian product formula own state * product of parent states
// 	ownState *= parentStates

// 	return ownState
// }

// func (n *Node) ConditionalValid() bool {
// 	// find out own total states
// 	totalStates := n.TotalStates()
// 	ownConditionalStates := 0

// 	// sum all of own states
// 	for _, ps := range n.Conditional {
// 		ownConditionalStates += len(ps.Space)
// 	}

// 	// return size check
// 	return ownConditionalStates == totalStates
// }
