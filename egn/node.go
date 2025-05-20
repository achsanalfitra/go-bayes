package egn

import (
	"fmt"
	"maps"
)

const (
	ConditionalType = "conditional"
	MarginalType    = "marginal"
	JointType       = "joint"
)

type Node struct {
	Name        string
	Context     *ProbabilityContext
	Marginal    *ProbabilitySpace
	Conditional map[string]*ProbabilitySpace // format {parentCombinations: space}
	Joint       map[string]*ProbabilitySpace // format {factors: space} where factors are variables including itself
	Parents     map[string]*Node
	Children    map[string]*Node
}

func NewNode(context *ProbabilityContext, name string) (*Node, error) {

	if _, nodeExists := context.NodeName[name]; !nodeExists {
		context.NodeName[name] = struct{}{}
	} else {
		return nil, fmt.Errorf("node already exists in this context")
	}

	return &Node{
		Name:        name,
		Context:     context,
		Marginal:    NewProbabilitySpace(),
		Conditional: make(map[string]*ProbabilitySpace),
		Joint:       make(map[string]*ProbabilitySpace),
		Parents:     make(map[string]*Node),
		Children:    make(map[string]*Node),
	}, nil
}

func (n *Node) SetMarginalRoot(event string, probability float64) error {
	if len(n.Parents) > 0 {
		return fmt.Errorf("can't specify marginal since the node %s already has at least one parent", n.Name)
	}

	// encode marginal probability event as "A=a"
	eventMap := map[string]string{n.Name: event}
	encodedMarginal := EncodeEvents(eventMap)

	// add probability pair to node
	n.Marginal.AddPair(encodedMarginal, probability)

	// update probability event to context ledger
	n.UpdateState(encodedMarginal, MarginalType, nil)

	return nil
}

func (n *Node) SetMarginal(event string, probability float64) error {
	if len(n.Parents) > 0 {
		return fmt.Errorf("can't specify marginal since the node %s already has at least one parent", n.Name)
	}

	// check if joint exists
	// decode factors from conditional

	for factors := range n.Context.Joint {
		factorsMap := make(map[string]struct{})
		factorsMap = DecodeFactors(factors) // assign factors to factorsMap

		if _, jointExists := factorsMap[n.Name]; jointExists {
			return fmt.Errorf("can't specify marginal since the node %s is already included on %s joint probabilities", n.Name, factors)
		}
	}

	// encode marginal probability event as "A=a"
	eventMap := map[string]string{n.Name: event}
	encodedMarginal := EncodeEvents(eventMap)

	// add probability pair to node
	n.Marginal.AddPair(encodedMarginal, probability)

	// update probability event to context ledger
	n.UpdateState(encodedMarginal, MarginalType, nil)

	return nil
}

func (n *Node) CompleteMarg() {
	if n.marg.TotalProb() < 1 && len(n.marg.space) > 0 {
		encodedMarg := fmt.Sprintf("%s=%s", n.name, "_")
		n.marg.AddPair(encodedMarg, 1-n.marg.TotalProb())
		n.UpdateState(encodedMarg, "marginal", nil)
	}
}

func (n *Node) NormalizeMarg() {
	if n.marg.TotalProb() != 1 && len(n.marg.space) > 0 {
		n.marg.Normalize()
	}
}

func (n *Node) SetCond(event string, givenState map[string]string, prob float64) error {
	// Check if marginal and joint are already set
	factors := make(map[string]string)
	factors[n.name] = event
	maps.Copy(factors, givenState) // factors stores with format {nodeName : event}

	// Check if dependency node exists
	for nodeName := range factors {
		if _, nodeExists := n.context.NodeName[nodeName]; !nodeExists {
			return fmt.Errorf("the dependent node %s doesn't exist", nodeName)
		}
	}

	factorCombinations := n.encodeFactors(factors) // factors with format "A B C"
	jointCombinations := n.encodeJoint(factors)    // joint with format "A=a B=b C=c"

	// Check if joint map exists
	// Checking an empty map for key is fine in Golang
	_, jointExists := n.context.Joint[factorCombinations][jointCombinations]

	// Check if marginal map exists
	_, margExists := n.context.Marginal[n.name][event]

	if margExists && jointExists {
		return fmt.Errorf("you can't specify conditional probability since the node %s already has marginal and joint probability specified", n.name)
	}

	// Update event into conditional probility

	// Create new probability space if it doesn't exist
	encodedParents := n.encodeParents(givenState)

	if n.cond[encodedParents] == nil {
		n.cond[encodedParents] = NewProbabilitySpace()
	}

	// add probability pair to node
	n.cond[encodedParents].AddPair(n.encodeCond(event, givenState), prob)

	// update probability event to context ledger
	n.UpdateState(n.encodeCond(event, givenState), "conditional", &givenState)

	return nil
}

func (n *Node) CompleteCond() error {
	// check for conditional existence
	if len(n.context.Conditional) == 0 {
		return fmt.Errorf("no conditional probability exists for this node")
	}

	for parentCombinations := range n.context.Conditional {
		total := n.cond[parentCombinations].TotalProb()

		if total < 1 {
			decodedParents := n.decodeJoint(parentCombinations)
			encodedEvents := n.encodeCond("_", decodedParents)

			n.cond[parentCombinations].AddPair(encodedEvents, 1-total)
			n.UpdateState(encodedEvents, "conditional", &decodedParents)
		}
	}

	return nil
}

func (n *Node) NormalizeCond() error {
	// check for conditional existence
	if len(n.context.Conditional) == 0 {
		return fmt.Errorf("no conditional probability exists for this node")
	}

	for parentCombinations := range n.context.Conditional {
		total := n.cond[parentCombinations].TotalProb()

		if total != 1 {
			n.cond[parentCombinations].Normalize()
		}

	}
	return nil
}

func (n *Node) SetJoint(events map[string]string, prob float64) error {
	// Check if there is only one event listed
	if len(events) == 1 {
		return fmt.Errorf("joint events cannot be just one event")
	}

	// Check if dependency node exists
	for nodeName := range events {
		if _, nodeExists := n.context.NodeName[nodeName]; !nodeExists {
			return fmt.Errorf("the dependent node %s doesn't exist", nodeName)
		}
	}

	// Check if marginal map exists
	_, margExists := n.context.Marginal[n.name][events[n.name]]

	// Check if conditional exists
	parentCombinations := make(map[string]string)
	for parent, state := range events {
		if parent != n.name {
			parentCombinations[parent] = state
		}
	}

	encodedParents := n.encodeParents(parentCombinations)

	_, condExists := n.context.Conditional[n.name][encodedParents]

	if margExists && condExists {
		return fmt.Errorf("you can't specify conditional probability since the node %s already has marginal and conditional probability specified", n.name)
	}

	// Check if the node has conditional probability
	// Get parentCombination
	parentCombination := make(map[string]string)
	for parent, event := range events {
		if parent != n.name {
			parentCombination[parent] = event // Format {parentCombination: event}
		}
	}

	// Create new probability space if it doesn't exist
	encodedFactors := n.encodeFactors(events)
	jointEvents := n.encodeJoint(events)
	if n.joint[encodedFactors] == nil {
		n.joint[encodedFactors] = NewProbabilitySpace()
	}
	n.joint[encodedFactors].AddPair(jointEvents, prob)

	n.UpdateState(jointEvents, "joint", nil)

	return nil
}

func (n *Node) UpdateState(event string, probType string, parentState *map[string]string) {
	nodeName := n.name

	// Add state used in setting methods to the state list in the node
	switch probType {
	case "marg":
		// check outer map existence
		if _, ok := n.context.Marginal[nodeName]; !ok {
			n.context.Marginal[nodeName] = make(map[string]struct{})
		}

		if _, isExist := n.context.Marginal[nodeName][event]; !isExist {
			n.context.Marginal[nodeName][event] = struct{}{}
		}

	case "cond":
		if parentState != nil {
			encodedParents := n.encodeParents(*parentState)

			// check outer map existence
			if _, ok := n.context.Conditional[nodeName]; !ok {
				n.context.Conditional[nodeName] = make(map[string]map[string]struct{})
			}

			// check middle-layer map existence
			if _, ok := n.context.Conditional[nodeName][encodedParents]; !ok {
				n.context.Conditional[nodeName][encodedParents] = make(map[string]struct{})
			}

			// check the event existence, else add
			if _, isExist := n.context.Conditional[nodeName][encodedParents][event]; !isExist {
				n.context.Conditional[nodeName][encodedParents][event] = struct{}{}
			}
		}
	case "joint":
		decodedJoint := n.decodeJoint(event)
		encodedFactors := n.encodeFactors(decodedJoint)

		// check outer map existence
		if _, ok := n.context.Joint[encodedFactors]; !ok {
			n.context.Joint[encodedFactors] = make(map[string]struct{})
		}

		if _, isExist := n.context.Joint[encodedFactors][event]; !isExist {
			n.context.Joint[encodedFactors][event] = struct{}{}
		}
	}
}

func (n *Node) AddParent(parent *Node) {
	// Check if the parent already exists in the parents map
	if _, exists := n.parents[parent.name]; exists {
		fmt.Println("Parent node", parent.name, "already added.")
		return
	}

	// Add the parent to the node's parent map
	n.parents[parent.name] = parent
	// Optionally, also add this node as a child of the parent
	if n.children == nil {
		n.children = make(map[string]*Node)
	}
	n.children[n.name] = n

	fmt.Println("Added parent", parent.name, "to node", n.name)
}
