package engine

import "fmt"

// ProbabilitySpace is defined as pairs of sample and its probability
type ProbabilitySpace struct {
	space              map[string]float64
	isProbabilitySpace bool // Indicates whether the user current probability space is valid
}

// Create a new ProbabilitySpace
func NewProbabilitySpace() *ProbabilitySpace {
	return &ProbabilitySpace{
		space:              make(map[string]float64),
		isProbabilitySpace: true,
	}
}

// Add a new event and probability pair
func (ps *ProbabilitySpace) AddPair(event string, probability float64) {

	// Input first-check
	if probability < 0 || probability > 1 {
		fmt.Println("Error: enter probability value between 0 and 1")
		return
	}

	// Check if event already exist
	_, isExist := ps.space[event]

	if isExist {
		fmt.Println("Error: event already exists, use ChangeProbability to change its probability")
		return
	}

	// Check input against total probability
	totalProb := ps.CheckProb() + probability

	if totalProb > 1 {
		fmt.Println("Warning: your input makes total probability exceeding 1")
		ps.isProbabilitySpace = false
	} else {
		ps.space[event] = probability
		ps.isProbabilitySpace = true
	}
}

// Check total probability
func (ps *ProbabilitySpace) CheckProb() float64 {
	totalProb := 0.0
	for _, prob := range ps.space {
		totalProb += prob
	}
	return totalProb
}

// Check probability probability space validity
func (ps *ProbabilitySpace) IsValid() bool {
	return ps.isProbabilitySpace
}

// Show current probability space
func (ps *ProbabilitySpace) ShowPair() {
	for event, prob := range ps.space {
		fmt.Printf("%s: %.4f\n", event, prob)
	}
}

// Normalize each event probability
func (ps *ProbabilitySpace) Normalize() {
	totalProb := ps.CheckProb()

	if totalProb == 0 {
		fmt.Println("Error: total probability is zero, cannot normalize")
		return
	}

	for event, prob := range ps.space {
		ps.space[event] = prob / totalProb
	}

	ps.isProbabilitySpace = true
}
