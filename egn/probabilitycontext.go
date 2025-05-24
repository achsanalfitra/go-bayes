package egn

type ProbabilityContext struct {
	Marginal           map[string]map[string]struct{}                       // format {nodeName: {event : }}
	Conditional        map[string]map[string]map[string]struct{}            // format {nodeName: {dependencyCombination: {event: }}}
	PartialConditional map[string]map[string]map[string]map[string]struct{} // format {fullJointFactors {partialFactors: {dependencyCombinations: {event: }}}}
	Joint              map[string]map[string]struct{}                       // format {factors : {jointEvent: } }
	PartialJoint       map[string]map[string]map[string]struct{}            // format {fullJointFactors: {partialJointFactors: {event: }}} (partials must have full joint)
	NodeName           map[string]*Node                                     // format {nodeName: }
}

func BuildContext() *ProbabilityContext {
	return &ProbabilityContext{
		Marginal:           make(map[string]map[string]struct{}),
		Conditional:        make(map[string]map[string]map[string]struct{}),
		PartialConditional: make(map[string]map[string]map[string]map[string]struct{}),
		Joint:              make(map[string]map[string]struct{}),
		PartialJoint:       make(map[string]map[string]map[string]struct{}),
		NodeName:           make(map[string]*Node),
	}
}

func (pc *ProbabilityContext) ShowMarginal() {
	for nodeName := range pc.Marginal {
		pc.NodeName[nodeName].Show.MarginalEvents()
	}
}

func (pc *ProbabilityContext) CheckConsistency(nodeName string) {
	// Check itself on joint ledger, if exists check marginal event existence on the ledger
	// If exists check its marginal probability
	//
}
