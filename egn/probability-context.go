package egn

type ProbabilityContext struct {
	Marginal           map[string]map[string]struct{}            // format {nodeName: {event : }}
	Conditional        map[string]map[string]map[string]struct{} // format {nodeName: {dependencyCombination: {event: }}}
	PartialConditional map[string]map[string]map[string]struct{} // format {partialFactors: {dependencyCombinations: {event: }}}
	Joint              map[string]map[string]struct{}            // format {factors : {jointEvent: } }
	PartialJoint       map[string]map[string]map[string]struct{} // format {fullJointFactors: {partialJointFactors: {event: }}} (partials must have full joint)
	NodeName           map[string]struct{}                       // format {nodeName: }
}

func BuildContext() *ProbabilityContext {
	return &ProbabilityContext{
		Marginal:           make(map[string]map[string]struct{}),
		Conditional:        make(map[string]map[string]map[string]struct{}),
		PartialConditional: make(map[string]map[string]map[string]struct{}),
		Joint:              make(map[string]map[string]struct{}),
		PartialJoint:       make(map[string]map[string]map[string]struct{}),
		NodeName:           make(map[string]struct{}),
	}
}

// Conditional probability
