package egn

import (
	"fmt"
	"log"
	"maps"
)

type Set interface {
	NodeStates(states ...string) error
	MarginalProbability(event string, probability float64) error
	ConditionalProbability(event string, givenEvents map[string]string, probability float64) error
}

func (n *Node) NodeStates(states ...string) error {
	// create map of states
	statesMap := make(map[string]struct{})

	for _, state := range states {
		if _, stateExists := statesMap[state]; stateExists {
			return fmt.Errorf("your input state %s is duplicated", state)
		}

		// create states if duplicated
		statesMap[state] = struct{}{}
	}

	maps.Copy(n.States, statesMap)

	return nil
}

func (n *Node) MarginalProbability(event string, probability float64) error {
	// encode marginal probability event as "A=a"
	eventMap := map[string]string{n.Name: event}
	encodedMarginal := EncodeEvents(eventMap)

	// add probability pair to node
	err := n.Marginal.AddPair(encodedMarginal, probability)
	if err != nil {
		log.Println(err)
		return err
	}

	// update probability event to context ledger
	n.UpdateState(encodedMarginal, MarginalType, nil)

	return nil
}

func (n *Node) ConditionalProbability(event string, givenEvents map[string]string, probability float64) error {
	for name := range givenEvents {
		if _, nodeExists := n.Context.NodeName[name]; !nodeExists {
			return fmt.Errorf("the node %s doesn't exist in this context", name)
		}
	}

	encodedGivenEvents := EncodeEvents(givenEvents)

	eventMap := map[string]string{n.Name: event}
	encodedConditionalEvent := EncodeConditional(eventMap, givenEvents)

	// check probability space existence
	if _, spaceExists := n.Conditional[encodedGivenEvents]; !spaceExists {
		n.Conditional[encodedGivenEvents] = NewProbabilitySpace()
	}

	// add probability pair to node
	n.Conditional[encodedGivenEvents].AddPair(encodedConditionalEvent, probability)

	// update probability event to context ledger
	n.UpdateState(encodedConditionalEvent, ConditionalType, &givenEvents)

	return nil
}
