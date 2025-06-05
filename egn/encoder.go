package egn

import (
	"fmt"
	"slices"
	"strings"

	"github.com/achsanalfitra/go-bayes/hlp"
)

func SingleEventToString(name, state string) string {
	var encoded strings.Builder

	encoded.WriteString(name)
	encoded.WriteString("=")
	encoded.WriteString(state)

	return encoded.String()
}

// for now, this function mimics the given events completely
func MultiEventToString(events map[string]string, eventsMap *hlp.BiMapInt) (string, error) {
	// no need to point to the node object, the map should suffice
	ordered := []int{}

	// retrieve order
	for parent := range events {
		if _, exist := eventsMap.StrInt[parent]; !exist {
			return "", fmt.Errorf("the given parent %s doesn't exist", parent)
		}

		ordered = append(ordered, eventsMap.StrInt[parent])
	}

	slices.Sort(ordered)

	// build string
	var output strings.Builder

	// for each ordered int, get the respective parent from the given map
	for i := range ordered {
		parentKey := eventsMap.IntStr[ordered[i]]
		output.WriteString(parentKey)
		output.WriteString("=")
		output.WriteString(events[parentKey])
		if i != len(ordered)-1 {
			output.WriteString(" ")
		}
	}

	return output.String(), nil
}

func SingleEventToMap(event string) map[string]string {
	output := make(map[string]string)

	space := strings.Split(event, "=")

	output[space[0]] = space[1]

	return output
}

func MultiEventToMap(events string) map[string]string {
	output := make(map[string]string)

	eventsArray := strings.Split(events, " ")

	for _, event := range eventsArray {
		singleMap := SingleEventToMap(event)
		for name, state := range singleMap {
			output[name] = state
		}
	}

	return output
}

func GivenEventToString(events map[string]string, givenMap *hlp.BiMapInt) (string, error) {
	// no need to point to the node object, the map should suffice
	ordered := []int{}

	// retrieve order
	for parent := range events {
		if _, exist := givenMap.StrInt[parent]; !exist {
			return "", fmt.Errorf("the given parent %s doesn't exist", parent)
		}

		ordered = append(ordered, givenMap.StrInt[parent])
	}

	slices.Sort(ordered)

	// build string
	var output strings.Builder

	// for each ordered int, get the respective parent from the given map
	for i := range ordered {
		parentKey := givenMap.IntStr[ordered[i]]
		output.WriteString(parentKey)
		output.WriteString("=")
		output.WriteString(events[parentKey])
		if i != len(ordered)-1 {
			output.WriteString(" ")
		}
	}

	return output.String(), nil
}

func ConditionalToString(name, state string, givenEvents map[string]string, givenMap *hlp.BiMapInt) (string, error) {
	var output strings.Builder

	// write name and state
	eventStr := SingleEventToString(name, state)
	output.WriteString(eventStr)

	// write the pipe separator
	output.WriteString(" | ")

	// build given events string
	givenEventsStr, err := GivenEventToString(givenEvents, givenMap)
	if err != nil {
		return "", err
	}

	output.WriteString(givenEventsStr)

	return output.String(), nil
}

func ConditionalToMap(conditionalEvent string) (map[string]string, map[string]string) {
	eventOutput := make(map[string]string)
	givenEventOutput := make(map[string]string)

	// split to event and givenEvent
	conditionalEventArray := strings.Split(conditionalEvent, " | ")
	event := conditionalEventArray[0]
	givenEvent := conditionalEventArray[1]

	// assign event to map
	eventOutput = SingleEventToMap(event)

	// assign multi events to map
	givenEventOutput = MultiEventToMap(givenEvent)

	return eventOutput, givenEventOutput
}
