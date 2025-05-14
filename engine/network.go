package engine

type Network struct {
	context *ProbabilityContext
	name    string
	nodes   map[string]*Node           // Format {name: node object}
	edges   map[string]map[string]bool // Format {name: {child:true}}

}

func NewNetwork(context *ProbabilityContext, name string) *Network {
	return &Network{
		context: context,
		name:    name,
		nodes:   make(map[string]*Node),
		edges:   make(map[string]map[string]bool),
	}
}

func (n *Network) AddNode(node *Node) {
	n.nodes[node.name] = node
}

func (n *Network) AddEdge(parent string, child string) {
	if _, parentExists := n.nodes[parent]; !parentExists {
		return
	}

	if _, childExists := n.nodes[child]; !childExists {
		return
	}

	if _, isExist := n.edges[parent]; !isExist {
		n.edges[parent] = make(map[string]bool)
	}

	if _, isChild := n.edges[parent][child]; !isChild {
		n.edges[parent][child] = true
		n.nodes[child].AddParent(n.nodes[parent])
	}
}
