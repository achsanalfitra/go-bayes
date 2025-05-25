package egn

import "fmt"

type Validator interface {
	CheckNode()         //	individual check
	CheckMarginal()     // check marginal consistency
	CheckConditional()  // check conditional consistency
	CheckJoint()        // check the existence of complete joint
	CheckCompleteness() // check the completeness of all
	CheckInferrable()   // check whether the context is ready for inference or not
}

func (pc *ProbabilityContext) CheckNode(nodeName string) (bool, error) {
	if _, nodeExists := pc.NodeName[nodeName]; !nodeExists {
		return false, fmt.Errorf("Node doesn't exist")
	}

	node := pc.NodeName[nodeName]

}
