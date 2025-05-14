package egn

type ProbabilityContext struct {
	Marginal    map[string]map[string]float64            // format {nodeName: {event : value}}
	Conditional map[string]map[string]map[string]float64 // format {nodeName: {parentCombination: {event: value}}}
	Joint       map[string]float64                       // format {jointEvent : value}
}

func BuildContext() *ProbabilityContext {
	return &ProbabilityContext{
		Marginal:    make(map[string]map[string]float64),
		Conditional: make(map[string]map[string]map[string]float64),
		Joint:       make(map[string]float64),
	}
}

// Conditional probability
