package egn

import "fmt"

// ProbabilitySpace is defined as pairs of sample and its probability
type ProbabilitySpace struct {
	Space   map[string]float64
	IsValid bool // Indicates whether the user current probability space is valid
}

// Create a new ProbabilitySpace
func NewProbabilitySpace() *ProbabilitySpace {
	return &ProbabilitySpace{
		Space:   make(map[string]float64),
		IsValid: true,
	}
}

// Add a new event and probability pair
func (ps *ProbabilitySpace) AddPair(event string, probability float64) error {

	// Input first-check
	if probability < 0 || probability > 1 {
		return fmt.Errorf("enter probability value between 0 and 1")
	}

	// Check if event already exist
	_, isExist := ps.Space[event]

	if isExist {
		return fmt.Errorf("event %s already exists, use ChangeProbability to change its probability", event)
	}

	// Add the input pair
	ps.Space[event] = probability
	ps.UpdateValidity()

	// Warning the user if the probability exceeds 1
	if !ps.CheckValidity() {
		fmt.Println("Warning: your input makes total probability exceeding 1")
	}

	return nil
}

// Remove a pair of sample
func (ps *ProbabilitySpace) DeletePair(event string) {
	// Check if event already exist
	_, isExist := ps.Space[event]

	if !isExist {
		fmt.Println("Error: event doesn't exist")
		return
	}

	delete(ps.Space, event)
	ps.UpdateValidity()
}

// Change probability of an event
func (ps *ProbabilitySpace) UpdateProbability(event string, prob float64) {

	// Input first-check
	if prob < 0 || prob > 1 {
		fmt.Println("Error: enter probability value between 0 and 1")
		return
	}

	// Check if event already exist
	_, isExist := ps.Space[event]

	if isExist {
		ps.Space[event] = prob
		ps.UpdateValidity()
	} else {
		fmt.Println("Error: event doesn't exist")
		return
	}
}

// Check total probability
func (ps *ProbabilitySpace) TotalProb() float64 {
	totalProb := 0.0
	for _, prob := range ps.Space {
		totalProb += prob
	}
	return totalProb
}

// Check probability probability space validity
func (ps *ProbabilitySpace) CheckValidity() bool {
	return ps.IsValid
}

// Update probability space validity
func (ps *ProbabilitySpace) UpdateValidity() {
	totalProb := ps.TotalProb()

	ps.IsValid = totalProb >= 0 && totalProb <= 1
}

// Show current probability space
func (ps *ProbabilitySpace) ShowPair() {
	for event, prob := range ps.Space {
		fmt.Printf("%s: %.4f\n", event, prob)
	}
}

// Normalize each event probability
func (ps *ProbabilitySpace) Normalize() {
	totalProb := ps.TotalProb()

	if totalProb == 0 {
		fmt.Println("Error: total probability is zero, cannot normalize")
		return
	}

	for event, prob := range ps.Space {

		ps.Space[event] = prob / totalProb
	}

	ps.IsValid = true
}
