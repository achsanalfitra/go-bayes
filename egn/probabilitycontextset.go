package egn

import "fmt"

type Define interface {
	JointProbability() error
}

func (pc *ProbabilityContext) JointProbability(events map[string]string) error {
	if len(events) != len(pc.NodeName) {
		return fmt.Errorf("you can only define joint if you include all the existing factors/nodes")
	}

	// implement add pair
	// implement update state to context

	return nil
}
