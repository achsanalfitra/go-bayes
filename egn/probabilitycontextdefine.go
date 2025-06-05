package egn

// import "fmt"

type Define interface {
	JointProbability(events map[string]string, probability float64) error
	JointFactors(factors map[string]struct{}) error
}

// func (pc *ProbabilityContext) JointProbability(events map[string]string, probability float64) error {
// 	// check whether factors exist in context node
// 	encodedFactors := EncodeFactorsFromEvents(events)

// 	if _, factorsExist := pc.Joint[encodedFactors]; !factorsExist {
// 		return fmt.Errorf("factors do not exists")
// 	}

// 	// add probability pair to node
// 	encodedEvents := EncodeEvents(events)

// 	pc.Joint[encodedFactors].AddPair(encodedEvents, probability)

// 	return nil
// }

// func (pc *ProbabilityContext) JointFactors(factors map[string]struct{}) error {
// 	// check whether factors exist in context node
// 	for name := range factors {
// 		if _, factorExists := pc.NodeName[name]; !factorExists {
// 			return fmt.Errorf("factors do not exists")
// 		}
// 	}

// 	encodedFactors := EncodeFactors(factors)

// 	if _, factorsExist := pc.Joint[encodedFactors]; !factorsExist {
// 		pc.Joint[encodedFactors] = NewProbabilitySpace()
// 	}

// 	return nil
// }
