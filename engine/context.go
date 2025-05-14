package engine

type ProbabilityContext struct {
	Marginal    map[string]map[string]float64 // format {nodeName: {event : value}}
	Conditional map[string]map[string]float64 // call each probability space to infer type
	Joint       map[string]map[string]float64
}

func BuildContext() *ProbabilityContext {
	return &ProbabilityContext{
		Marginal:    make(map[string]map[string]float64),
		Conditional: make(map[string]map[string]float64),
		Joint:       make(map[string]map[string]float64),
	}
}
