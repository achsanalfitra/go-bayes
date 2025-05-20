package egn

import (
	"fmt"
	"log"
)

type Set interface {
	MarginalProbability(event string, probability float64) error
	ConditionalProbability(event string, givenEvents map[string]string, probability float64) error
	JointProbability(events map[string]string, probability float64) error
}

func (n *Node) MarginalProbability(event string, probability float64) error {
	// check if parent exists
	if len(n.Parents) > 0 {
		log.Printf("the node %s has a parent/parents", n.Name)
		return fmt.Errorf("the node %s has a parent/parents", n.Name)
	}

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

func (n *Node) JointProbability(events map[string]string, probability float64) error {
	// Check if there is only one event listed
	if len(events) < 2 {
		return fmt.Errorf("joint events cannot be just one event")
	}

	// check if the nodes exist
	for name := range events {
		if _, nodeExists := n.Context.NodeName[name]; !nodeExists {
			return fmt.Errorf("the node %s doesn't exist in this context", name)
		}
	}

	// Create new probability space if it doesn't exist
	encodedFactors := EncodeFactors(events)
	jointEvents := EncodeEvents(events)
	if _, jointExists := n.Joint[encodedFactors]; !jointExists {
		n.Joint[encodedFactors] = NewProbabilitySpace()
	}
	n.Joint[encodedFactors].AddPair(jointEvents, probability)

	n.UpdateState(jointEvents, JointType, nil)

	return nil
}
