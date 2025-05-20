package egn

import (
	"slices"
	"strings"
)

func EncodeEvents(events map[string]string) string {
	var encoded strings.Builder

	// sort event names so it is deterministic
	eventNames := make([]string, 0, len(events))
	for eventName := range events {
		eventNames = append(eventNames, eventName)
	}

	slices.Sort(eventNames)

	// add sorted event state onto encoded in order
	for i, eventName := range eventNames {
		encoded.WriteString(eventName)
		encoded.WriteString("=")
		encoded.WriteString(events[eventName]) // map query returns event state
		if i != len(eventNames)-1 {
			encoded.WriteString(" ") // add whitespace only if not the last event
		}
	}

	return encoded.String() // output format "A=b B=c" also works with singular event
}

func EncodeConditional(events map[string]string, givenEvents map[string]string) string {
	// Create encoded as strings.Builder

	encodedEvents := EncodeEvents(events) // call helper function to deterministically sort then encode
	encodedGivenEvents := EncodeEvents(givenEvents)

	var encoded strings.Builder
	encoded.WriteString(encodedEvents)
	encoded.WriteString(" | ")
	encoded.WriteString(encodedGivenEvents)

	// Return encoded as string
	return encoded.String() // Output format "A=a | B=b C=c"
}

func EncodeFactors(factors map[string]string) string {
	var encoded strings.Builder

	// Sort factors so it is deterministic
	names := make([]string, 0, len(factors))
	for name := range factors {
		names = append(names, name)
	}

	slices.Sort(names)

	// Adding sorted parent state onto encoded in order
	for i, name := range names {
		encoded.WriteString(name)
		if i != len(names)-1 {
			encoded.WriteString(" ") // Add whitespace only if not the last event
		}
	}

	return encoded.String()
}

func (n *Node) decodeCond(encodedCond string) map[string]string {
	output := make(map[string]string)

	pipeRemoved := strings.Split(encodedCond, " | ") // current ["A=a", "B=b C=c"]
	eventPair := strings.Split(pipeRemoved[0], "=")  // eventPair = ["A", "a"]
	output[eventPair[0]] = eventPair[1]

	parentPairs := strings.Fields(pipeRemoved[1]) // parentPair = ["B=b", "C=c"], split on whitespace with strings.Fields()
	for _, parent := range parentPairs {          // parent = ["B=b"]
		pair := strings.Split(parent, "=") // pair = ["B", "b"]
		output[pair[0]] = pair[1]
	}

	return output
}

func (n *Node) decodeJoint(encodedJoint string) map[string]string {
	output := make(map[string]string)

	factorPairs := strings.Fields(encodedJoint) // factorPairs = ["A=a", "B=b"], split on whitespace with strings.Fields()
	for _, factor := range factorPairs {        // parent = ["A=a"]
		pair := strings.Split(factor, "=") // pair = ["A", "a"]
		output[pair[0]] = pair[1]
	}

	return output
}
