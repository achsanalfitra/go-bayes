package engine

import (
	"fmt"
	"slices"
	"strings"
)

type Node struct {
	context  *ProbabilityContext
	name     string
	marg     *ProbabilitySpace
	cond     *ProbabilitySpace
	joint    *ProbabilitySpace
	parents  map[string]*Node
	children map[string]*Node
	cpt      map[string]float64
}

func NewNode(context *ProbabilityContext, name string) *Node {
	return &Node{
		name:     name,
		context:  context,
		marg:     NewProbabilitySpace(),
		cond:     NewProbabilitySpace(),
		joint:    NewProbabilitySpace(),
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

func (n *Node) SetCond(event string, givenState map[string]string, prob float64) {
	// Check if parents do not exist
	if len(n.parents) == 0 {
		fmt.Println("Error: you can't specify conditional probability since the node", n.name, "has no parents")

		return
	}

	// Check if marginal and joint are already set
	localJointState := make(map[string]string)
	for name, state := range givenState {
		localJointState[name] = state
	}
	localJointState[n.name] = event

	_, margExist := n.context.Marginal[n.name][event]
	_, jointExist := n.context.Joint[n.name][n.encodeJoint(localJointState)]

	if margExist && jointExist {
		fmt.Println("Error: you can't specify conditional probability since the node", n.name, "already has marginal and joint probability specified")

		return
	}

	// Update event into conditional probility
	n.cond.AddPair(n.encodeCond(event, givenState), prob)

	// Update state into node states
	n.UpdateState(n.encodeCond(event, givenState), prob, "conditional", nil)
}

func (n *Node) SetJoint(prob float64, events map[string]string) {
	// Check if there is only one event listed
	if len(events) == 1 {
		fmt.Println("Cant have 1 event joint probability bro")
	}

	// Check if the node has been set as
	_, margExist := n.context.Marginal[n.name][events[n.name]]

	// Duplicate events and except the node event to get parents
	parents := make(map[string]string)
	for k, v := range events {
		if k != n.name {
			parents[k] = v
		}
	}
	_, condExist := n.context.Conditional[n.name][n.encodeCond(events[n.name], parents)]

	if margExist && condExist {
		fmt.Println("Error: you can't specify conditional probability since the node", n.name, "already has marginal and conditional probability specified")

		return
	}

	n.joint.AddPair(n.encodeJoint(events), prob)

	for name := range events {
		n.UpdateState(n.encodeJoint(events), prob, "joint", &name)
	}
}

func (n *Node) UpdateState(event string, prob float64, probType string, name *string) {
	if name == nil {
		name = &n.name
	}

	// Add state used in setting methods to the state list in the node
	switch probType {
	case "marg":
		if _, isExist := n.context.Marginal[*name][event]; !isExist {
			n.context.Marginal[*name][event] = prob
		}
	case "cond":
		if _, isExist := n.context.Conditional[*name][event]; !isExist {
			n.context.Conditional[*name][event] = prob
		}
	case "joint":
		if _, isExist := n.context.Joint[*name][event]; !isExist {
			n.context.Joint[*name][event] = prob
		}
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
