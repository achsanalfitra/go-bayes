package egn

import (
	"fmt"
)

const (
	ConditionalType = "conditional"
	MarginalType    = "marginal"
	JointType       = "joint"
)

type Node struct {
	Name        string
	Context     *ProbabilityContext
	States      map[string]struct{}
	Marginal    *ProbabilitySpace
	Conditional map[string]*ProbabilitySpace // format {parentStateCombinations: space}
	Parents     map[string]*Node             // a node acknowledges who are its parents
	Children    map[string]*Node             // and so does its children
	Set         Set
	Show        Show
}

func NewNode(context *ProbabilityContext, name string) (*Node, error) {

	// check node existence
	if _, nodeExists := context.NodeName[name]; !nodeExists {
		node := &Node{
			Name:        name,
			Context:     context,
			States:      make(map[string]struct{}),
			Marginal:    NewProbabilitySpace(),
			Conditional: make(map[string]*ProbabilitySpace),
			Parents:     make(map[string]*Node),
			Children:    make(map[string]*Node),
		}

		context.NodeName[name] = node

		node.Set = node //set interface to itself to initialize it
		node.Show = node

		return node, nil
	}

	return nil, fmt.Errorf("node already exists in this context")
}

func (n *Node) UpdateState(encodedEvent string, probType string, givenEvents *map[string]string) {
	// input format
	// event = encodedEvents any type
	// givenEvents = map

	nodeName := n.Name

	// Add state used in setting methods to the state list in the node
	switch probType {
	case "marginal":
		// check outer map existence
		if _, ok := n.Context.Marginal[nodeName]; !ok {
			n.Context.Marginal[nodeName] = make(map[string]struct{})
		}

		if _, isExist := n.Context.Marginal[nodeName][encodedEvent]; !isExist {
			n.Context.Marginal[nodeName][encodedEvent] = struct{}{}
		}

	case "conditional":
		if givenEvents != nil {
			encodedGivenEvents := EncodeEvents(*givenEvents)

			// check outer map existence
			if _, ok := n.Context.Conditional[nodeName]; !ok {
				n.Context.Conditional[nodeName] = make(map[string]map[string]struct{})
			}

			// check middle-layer map existence
			if _, ok := n.Context.Conditional[nodeName][encodedGivenEvents]; !ok {
				n.Context.Conditional[nodeName][encodedGivenEvents] = make(map[string]struct{})
			}

			// check the event existence, else add
			if _, isExist := n.Context.Conditional[nodeName][encodedGivenEvents][encodedEvent]; !isExist {
				n.Context.Conditional[nodeName][encodedGivenEvents][encodedEvent] = struct{}{}
			}
		}
	}
}

func (n *Node) AddParent(parent *Node) error {
	// check if the parent already exists in the parents map
	if _, parentExists := n.Parents[parent.Name]; parentExists {
		fmt.Println("Parent node", parent.Name, "already added.")
		return fmt.Errorf("parent node %s is already added", parent.Name)
	}

	// add the parent to the node's parent map
	n.Parents[parent.Name] = parent

	// add the node to the children
	parent.Children[n.Name] = n

	return nil
}

// func (n *Node) CompleteMarg() {
// 	if n.marg.TotalProb() < 1 && len(n.marg.space) > 0 {
// 		encodedMarg := fmt.Sprintf("%s=%s", n.name, "_")
// 		n.marg.AddPair(encodedMarg, 1-n.marg.TotalProb())
// 		n.UpdateState(encodedMarg, "marginal", nil)
// 	}
// }

// func (n *Node) NormalizeMarg() {
// 	if n.marg.TotalProb() != 1 && len(n.marg.space) > 0 {
// 		n.marg.Normalize()
// 	}
// }

// func (n *Node) CompleteCond() error {
// 	// check for conditional existence
// 	if len(n.context.Conditional) == 0 {
// 		return fmt.Errorf("no conditional probability exists for this node")
// 	}

// 	for parentCombinations := range n.context.Conditional {
// 		total := n.cond[parentCombinations].TotalProb()

// 		if total < 1 {
// 			decodedParents := n.decodeJoint(parentCombinations)
// 			encodedEvents := n.encodeCond("_", decodedParents)

// 			n.cond[parentCombinations].AddPair(encodedEvents, 1-total)
// 			n.UpdateState(encodedEvents, "conditional", &decodedParents)
// 		}
// 	}

// 	return nil
// }

// func (n *Node) NormalizeCond() error {
// 	// check for conditional existence
// 	if len(n.context.Conditional) == 0 {
// 		return fmt.Errorf("no conditional probability exists for this node")
// 	}

// 	for parentCombinations := range n.context.Conditional {
// 		total := n.cond[parentCombinations].TotalProb()

// 		if total != 1 {
// 			n.cond[parentCombinations].Normalize()
// 		}

// 	}
// 	return nil
// }
