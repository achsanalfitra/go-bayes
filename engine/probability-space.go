package engine

import "fmt"

// ProbabilitySpace is defined as pairs of sample and its probability
type ProbabilitySpace struct {
	space map[string]float64
}

// Create a new ProbabilitySpace
func NewProbabilitySpace() *ProbabilitySpace {
	return &ProbabilitySpace{
		space: make(map[string]float64),
	}
}

// Add a new event and probability pair
func (ps *ProbabilitySpace) AddPair(event string, probability float64) {
	if probability < 0 || probability > 1 {
		fmt.Println("Error: enter probability value between 0 and 1")
		return
	}
	ps.space[event] = probability
}
