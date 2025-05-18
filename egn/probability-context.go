package egn

type ProbabilityContext struct {
	Marginal    map[string]map[string]struct{}            // format {nodeName: {event : }}
	Conditional map[string]map[string]map[string]struct{} // format {nodeName: {parentCombination: {event: }}}
	Joint       map[string]map[string]struct{}            // format {factors : {jointEvent: } }
	NodeName    map[string]struct{}                       // format {nodeName: }
}

func BuildContext() *ProbabilityContext {
	return &ProbabilityContext{
		Marginal:    make(map[string]map[string]struct{}),
		Conditional: make(map[string]map[string]map[string]struct{}),
		Joint:       make(map[string]map[string]struct{}),
		NodeName:    make(map[string]struct{}),
	}
}

// Conditional probability
