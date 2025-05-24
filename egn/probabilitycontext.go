package egn

type ProbabilityContext struct {
	Marginal    map[string]map[string]struct{}            // format {nodeName: {event : }}
	Conditional map[string]map[string]map[string]struct{} // format {nodeName: {dependencyCombination: {event: }}}
	Joint       map[string]*ProbabilitySpace              // format {factors : {space: {eventspace: value}} }
	NodeName    map[string]*Node                          // format {nodeName: }
}

func BuildContext() *ProbabilityContext {
	return &ProbabilityContext{
		Marginal:    make(map[string]map[string]struct{}),
		Conditional: make(map[string]map[string]map[string]struct{}),
		Joint:       make(map[string]*ProbabilitySpace),
		NodeName:    make(map[string]*Node),
	}
}

func (pc *ProbabilityContext) ShowMarginal() {
	for nodeName := range pc.Marginal {
		pc.NodeName[nodeName].Show.MarginalEvents()
	}
}
