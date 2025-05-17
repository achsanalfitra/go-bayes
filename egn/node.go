package egn

import (
	"fmt"
	"slices"
	"strings"
)

const (
	Cond  = "conditional"
	Marg  = "marginal"
	Joint = "joint"
)

type Node struct {
	context  *ProbabilityContext
	name     string
	marg     *ProbabilitySpace
	cond     map[string]*ProbabilitySpace // format {parentCombinations: space}
	joint    map[string]*ProbabilitySpace // format {factors: space} where factors are variables including itself
	parents  map[string]*Node
	children map[string]*Node
	cpt      map[string]float64
}

func NewNode(context *ProbabilityContext, name string) *Node {
	return &Node{
		name:     name,
		context:  context,
		marg:     NewProbabilitySpace(),
		cond:     make(map[string]*ProbabilitySpace),
		joint:    make(map[string]*ProbabilitySpace),
		parents:  make(map[string]*Node),
		children: make(map[string]*Node),
		cpt:      make(map[string]float64),
	}
}

func (n *Node) SetMarg(event string, prob float64) {
	if len(n.parents) > 0 {
		fmt.Println("Error: you can't specify marginal probability since the node", n.name, "these parents")
		for _, parent := range n.parents {
			fmt.Print(parent.name, " ")
		}
		fmt.Printf("\n")

		return
	}
	encodedMarg := fmt.Sprintf("%s=%s", n.name, event)
	n.marg.AddPair(encodedMarg, prob)
	n.UpdateState(encodedMarg, prob, "marginal", nil)
}

func (n *Node) CompleteMarg() {
	if n.marg.TotalProb() < 1 && len(n.marg.space) > 0 {
		encodedMarg := fmt.Sprintf("%s=%s", n.name, "_")
		n.marg.AddPair(encodedMarg, 1-n.marg.TotalProb())
		n.UpdateState(encodedMarg, 1-n.marg.TotalProb(), "marginal", nil)
	}
}

func (n *Node) NormalizeMarg() {
	if n.marg.TotalProb() != 1 && len(n.marg.space) > 0 {
		n.marg.Normalize()
	}
}

func (n *Node) SetCond(event string, givenState map[string]string, prob float64) {
	// Check if parents do not exist
	if len(n.parents) == 0 {
		fmt.Println("Error: you can't specify conditional probability since the node", n.name, "has no parents")

		return
	}

	// Check if marginal and joint are already set
	factors := make(map[string]string)
	factors[n.name] = event
	for name, state := range givenState {
		factors[name] = state
	} // factors stores with format {nodeName : event}

	// Check if joint map exists
	jointFactorsMap, jointOk := n.context.Joint[n.encodeFactors(factors)]
	jointExists := jointOk && jointFactorsMap != nil && jointFactorsMap[n.encodeJoint(factors)] != struct{}{}

	// Check if marginal map exists
	margMap, margOk := n.context.Marginal[n.name]
	margExists := margOk && margMap[event] != struct{}{}

	if margExists && jointExists {
		fmt.Println("Error: you can't specify conditional probability since the node", n.name, "already has marginal and joint probability specified")

		return
	}

	// Update event into conditional probility

	// Create new probability space if it doesn't exist
	key := n.encodeParents(givenState)
	if n.cond[key] == nil {
		n.cond[key] = NewProbabilitySpace()
	}
	n.cond[n.encodeFactors(factors)].AddPair(n.encodeCond(event, givenState), prob)

	// Update state into node states
	n.UpdateState(n.encodeCond(event, givenState), prob, "conditional", &givenState)
}

// func (n *Node) CompleteCond() {
// 	if n.cond.TotalProb() < 1 && len(n.cond.space) > 0 {
// 		var pNames []string

// 		for _, parent := range n.parents {
// 			pNames = append(pNames, parent.name)
// 		}

// 		for _, pName := range pNames {
// 			if len(n.parents[pName].parents) > 0 {
// 				// Insert adding complement condition for each parent combinations
// 			} else {
// 				n.cond.AddPair(encodedMarg, 1-n.marg.TotalProb())
// 				n.UpdateState(encodedMarg, 1-n.marg.TotalProb(), "marginal", nil)
// 			}
// 		}
// 	}

// 	if n.cond.TotalProb() < 1 && len(n.cond.space) > 0 {
// 		parentStates := make(map[string]map[string]bool)
// 		givenStates := make(map[string]map[string]string)

// 		for _, parent := range n.parents {
// 			for state, _ := range n.parents[parent.name].states {
// 				parentStates[parent.name][state] = true
// 			}

// 		}

// 	for _, parent := range n.parents {
// 		for _, innerParent := range

// 		n.cond.AddPair(n.encodeCond("_", givenState), 1-n.cond.TotalProb())
// 		n.UpdateState(n.encodeCond("_", givenState), 1-n.cond.TotalProb(), "conditional", nil)
// 	}
// }
// 	}
// }

// func (n *Node) NormalizeCond() {
// 	if n.cond.TotalProb() != 1 && len(n.cond.space) > 0 {
// 		n.cond.Normalize()
// 	}
// }

func (n *Node) SetJoint(events map[string]string, prob float64) {
	// Check if there is only one event listed
	if len(events) == 1 {
		fmt.Println("Cant have 1 event joint probability bro")
	}

	// Check marginal inner map
	if _, ok := n.context.Marginal[n.name]; !ok {
		n.context.Marginal[n.name] = make(map[string]struct{})
	}

	margMap, margOk := n.context.Marginal[n.name]
	_, margEventOk := margMap[events[n.name]]
	margExists := margOk && margEventOk

	// Check if the node has conditional probability
	// Get parentCombination
	parentCombination := make(map[string]string)
	for parent, event := range events {
		if parent != n.name {
			parentCombination[parent] = event // Format {parentCombination: event}
		}
	}

	parentCombinations := n.encodeParents(parentCombination) // Format "parentCombinations"

	// Ensure 2nd-level map exists for nodeName
	if _, ok := n.context.Conditional[n.name]; !ok {
		n.context.Conditional[n.name] = make(map[string]map[string]struct{})
	}

	// Ensure 3rd-level map exists for parentKey
	if _, ok := n.context.Conditional[n.name][parentCombinations]; !ok {
		n.context.Conditional[n.name][parentCombinations] = make(map[string]struct{})
	}

	condMap, condOk := n.context.Conditional[n.name][parentCombinations]
	_, condEventOk := condMap[events[n.name]]
	condExists := condOk && condEventOk

	if margExists && condExists {
		fmt.Println("Error: you can't specify conditional probability since the node", n.name, "already has marginal and conditional probability specified")

		return
	}

	// Create new probability space if it doesn't exist
	key := n.encodeFactors(events)
	jointEvents := n.encodeJoint(events)
	if n.joint[key] == nil {
		n.joint[key] = NewProbabilitySpace()
	}
	n.joint[key].AddPair(jointEvents, prob)

	n.UpdateState(jointEvents, prob, "joint", nil)
}

func (n *Node) UpdateState(event string, prob float64, probType string, parentState *map[string]string) {

	// Add state used in setting methods to the state list in the node
	switch probType {
	case "marg":
		// Ensure inner map exists
		if _, ok := n.context.Marginal[n.name]; !ok {
			n.context.Marginal[n.name] = make(map[string]struct{})
		}

		n.context.Marginal[n.name][event] = struct{}{}

	case "cond":
		if parentState != nil {
			nodeName := n.name
			parentKey := n.encodeParents(*parentState)
			eventKey := event

			// Ensure 2nd-level map exists for nodeName
			if _, ok := n.context.Conditional[nodeName]; !ok {
				n.context.Conditional[nodeName] = make(map[string]map[string]struct{})
			}

			// Ensure 3rd-level map exists for parentKey
			if _, ok := n.context.Conditional[nodeName][parentKey]; !ok {
				n.context.Conditional[nodeName][parentKey] = make(map[string]struct{})
			}

			// Finally, add the event key
			n.context.Conditional[nodeName][parentKey][eventKey] = struct{}{}
		}
	case "joint":
		// if _, isExist := n.context.Joint[event]; !isExist {
		// n.context.Joint[event] = prob
		// }
	}
}

func (n *Node) encodeCond(event string, parents map[string]string) string {
	// Create encoded as strings.Builder
	var encoded strings.Builder
	encoded.WriteString(n.name)
	encoded.WriteString("=")
	encoded.WriteString(event)
	encoded.WriteString(" |")

	// Sort parents so it is deterministic
	pNames := make([]string, 0, len(parents))
	for pName := range parents {
		pNames = append(pNames, pName)
	}

	slices.Sort(pNames)

	// Adding sorted parent state onto encoded in order
	for _, pName := range pNames {
		encoded.WriteString(" ")
		encoded.WriteString(pName)
		encoded.WriteString("=")
		encoded.WriteString(parents[pName]) // The map query returns pState
	}

	// Return encoded as string
	return encoded.String()
}

func (n *Node) encodeParents(parents map[string]string) string {
	var encoded strings.Builder

	// Sort parents so it is deterministic
	pNames := make([]string, 0, len(parents))
	for pName := range parents {
		pNames = append(pNames, pName)
	}

	slices.Sort(pNames)

	// Adding sorted parent state onto encoded in order
	for i, pName := range pNames {
		encoded.WriteString(pName)
		encoded.WriteString("=")
		encoded.WriteString(parents[pName]) // The map query returns pState
		if i != len(pNames)-1 {
			encoded.WriteString(" ") // Add whitespace only if not the last event
		}
	}

	return encoded.String()
}

func (n *Node) encodeFactors(factors map[string]string) string {
	var encoded strings.Builder

	// Sort factors so it is deterministic
	names := make([]string, 0, len(factors))
	for name := range factors {
		names = append(names, name)
	}

	slices.Sort(names)

	// Adding sorted parent state onto encoded in order
	for i, name := range names {
		encoded.WriteString(name)
		if i != len(names)-1 {
			encoded.WriteString(" ") // Add whitespace only if not the last event
		}
	}

	return encoded.String()
}

func (n *Node) encodeJoint(events map[string]string) string {
	// Create encoded as strings.Builder
	var encoded strings.Builder

	// Sort events so it is deterministic
	eNames := make([]string, 0, len(events))
	for eName := range events {
		eNames = append(eNames, eName)
	}

	slices.Sort(eNames)

	// Adding sorted event names onto encoded in order
	for i, eName := range eNames {
		encoded.WriteString(eName)
		encoded.WriteString("=")
		encoded.WriteString(events[eName])
		if i != len(eNames)-1 {
			encoded.WriteString(" ") // Add whitespace only if not the last event
		}
	}

	// Return encoded as string
	return encoded.String()
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
