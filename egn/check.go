package egn

func CanSetConditional(n *Node, event string, givenEvents map[string]string) bool {
	// check dividend and divisor existence

	dividend := make(map[string]string)

	for name, state := range givenEvents {
		dividend[name] = state
	}

	encodedDivisorEvents := EncodeEvents(dividend)
	encodedDivisorFactors := EncodeFactors(dividend)

	dividend[n.Name] = event

	encodedDividendEvents := EncodeEvents(dividend)
	encodedDividendFactors := EncodeFactors(dividend)

	var dividendExists bool

	if _, _Exists := n.Context.Joint[encodedDividendFactors][encodedDividendEvents]; _Exists {
		dividendExists = true
	}

	if dividendExists {
		if _, divisorExists := n.Context.Marginal[n.Name][encodedDivisorEvents]; divisorExists {
			return false
		}
		if _, divisorExists := n.Context.PartialJoint[encodedDividendFactors][encodedDivisorFactors]; divisorExists {
			return false
		}
	}

	return true
}

func CanComputeJoint(n *Node, events map[string]string) bool {
	// check if it has any partial
	encodedEvents := EncodeEvents(events)
	encodedEventsFactors := EncodeFactors(events)
	var partialFound bool
	var partialConditionalFound bool

	if _, partialExists := n.Context.PartialJoint[encodedEventsFactors]; partialExists {
		partialFound = true
	}

	if partialFound {
		if _, partialConditionalExists := n.Context.PartialConditional[encodedEventsFactors]; partialConditionalExists {
			partialConditionalFound = true
		}
	}

	if partialConditionalFound {
		for partialConditional := range n.Context.PartialConditional[encodedEventsFactors] {
			for partialConditionalEvent := range n.Context.PartialConditional[encodedEventsFactors][partialConditional] {
				decodedPartialConditionalEvents, decodedPartialConditionalGivenEvents := DecodeConditional(partialConditionalEvent)
				for name, state := range decodedPartialConditionalGivenEvents {
					decodedPartialConditionalEvents[name] = state
				}
				encodedPartialConditionalEvent := EncodeEvents(decodedPartialConditionalEvents)
				if encodedPartialConditionalEvent == encodedEvents {
					encodedPartialConditionalGivenEvent := EncodeEvents(decodedPartialConditionalGivenEvents)
					encodedPartialConditionalGivenFactors := EncodeFactors(decodedPartialConditionalGivenEvents)
					if _, partialConditionalGivenEventExist := n.Context.PartialJoint[encodedEventsFactors][encodedPartialConditionalGivenFactors][encodedPartialConditionalGivenEvent]; partialConditionalGivenEventExist {
						return true
					}
				}
			}
		}
	}

	// Check if the marginal has the conditional

	return false
}
