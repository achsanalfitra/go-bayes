package egn

import (
	"slices"
	"strings"
)

func SingleEventToString(name, state string) string {
	var encoded strings.Builder

	encoded.WriteString(name)
	encoded.WriteString("=")
	encoded.WriteString(state)

	return encoded.String()
}

func JointEventToString() {
	// implement
}

func GivenEventToString(events map[string]string, n *Node) {
	// implement encode from map to given events string, e.g. A: a, B: b -> "B=b A=a" (given that B is actually the first parent)
}

func ConditionalToString(event map[string]string, givenEvents map[string]string, n *Node) {
	// implement encode full conditional to string
	// use givenEventToString
	GivenEventToString(givenEvents, n)

	// then build the string
}

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

func EncodeFactorsFromEvents(factors map[string]string) string {
	var encoded strings.Builder

	// sort factors so it is deterministic
	names := make([]string, 0, len(factors))
	for name := range factors {
		names = append(names, name)
	}

	slices.Sort(names)

	// add sorted parent state onto encoded in order
	for i, name := range names {
		encoded.WriteString(name)
		if i != len(names)-1 {
			encoded.WriteString(" ") // add whitespace only if not the last event
		}
	}

	return encoded.String() // output format "A B C"
}

func EncodeFactors(factors map[string]struct{}) string {
	var encoded strings.Builder

	// sort factors so it is deterministic
	names := make([]string, 0, len(factors))
	for name := range factors {
		names = append(names, name)
	}

	slices.Sort(names)

	// add sorted parent state onto encoded in order
	for i, name := range names {
		encoded.WriteString(name)
		if i != len(names)-1 {
			encoded.WriteString(" ") // add whitespace only if not the last event
		}
	}

	return encoded.String() // output format "A B C"
}

func DecodeEvents(encodedEvents string) map[string]string {
	output := make(map[string]string)

	eventPairs := strings.Fields(encodedEvents) // factorPairs = ["A=a", "B=b"], split on whitespace with strings.Fields()
	for _, event := range eventPairs {          // parent = ["A=a"]
		pair := strings.Split(event, "=") // pair = ["A", "a"]
		output[pair[0]] = pair[1]
	}

	return output // output format map{"A": "a"}
}

func DecodeConditional(encodedConditional string) (map[string]string, map[string]string) {

	pipeRemoved := strings.Split(encodedConditional, " | ") // current ["A=a", "B=b C=c"]
	events := DecodeEvents(pipeRemoved[0])                  // call helper to assign each events to map
	givenEvents := DecodeEvents(pipeRemoved[1])

	return events, givenEvents // output format map{"A": "a"} and map{"B": "b"}
}

func DecodeFactors(encodedFactors string) map[string]struct{} {
	output := make(map[string]struct{})

	factors := strings.Fields(encodedFactors) // factors = ["A", "B"], split on whitespace with strings.Fields()
	for _, factor := range factors {          // factor = "A"
		output[factor] = struct{}{}
	}

	return output // output format map{"A": }
}

func DecodeFactorsFromEvents(encodedEvents string) map[string]struct{} {
	output := make(map[string]struct{})

	events := strings.Fields(encodedEvents) // events = ["A=a", "B=b"]
	for _, event := range events {          // event = "A=a"
		pair := strings.Split(event, "=") // pair = ["A", "a"]
		output[pair[0]] = struct{}{}
	}

	return output // output format map{"A": }
}

func DecodeFactorsFromConditional(encodedConditional string) (map[string]struct{}, map[string]struct{}) {

	pipeRemoved := strings.Split(encodedConditional, " | ") // current ["A=a", "B=b C=c"]
	events := DecodeFactorsFromEvents(pipeRemoved[0])       // call helper to assign each events to map
	givenEvents := DecodeFactorsFromEvents(pipeRemoved[1])

	return events, givenEvents // output format map{"A": } and map{"B": }
}
