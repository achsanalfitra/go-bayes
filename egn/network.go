package egn

import "fmt"

type Network struct {
	context *ProbabilityContext
	name    string
	nodes   map[string]*Node           // Format {name: node object}
	edges   map[string]map[string]bool // Format {name: {child:true}}
	roots   map[string]*Node           // Format {name: node object}
}

func NewNetwork(context *ProbabilityContext, name string) *Network {
	return &Network{
		context: context,
		name:    name,
		nodes:   make(map[string]*Node),
		edges:   make(map[string]map[string]bool),
		roots:   make(map[string]*Node),
	}
}

func (n *Network) AddNode(node *Node) {
	n.nodes[node.name] = node
}

func (n *Network) AddEdge(parent *Node, child *Node) {
	if _, parentExists := n.nodes[parent.name]; !parentExists {
		return
	}

	if _, childExists := n.nodes[child.name]; !childExists {
		return
	}

	if _, isExist := n.edges[parent.name]; !isExist {
		n.edges[parent.name] = make(map[string]bool)
	}

	if _, isRoot := n.roots[child.name]; isRoot {
		delete(n.roots, child.name)
	}

	if _, isChild := n.edges[parent.name][child.name]; !isChild {
		n.edges[parent.name][child.name] = true
		n.nodes[child.name].AddParent(n.nodes[parent.name])
	}
}

func (n *Network) AddRoot(node *Node) {
	if len(node.parents) == 0 {
		n.roots[node.name] = node
	} else {
		fmt.Println("Error: this node has a parent")
		return
	}
}
