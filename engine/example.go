package engine

import "fmt"

func Test() {
	context := BuildContext()

	network := NewNetwork(context, "Test")

	n1 := NewNode(context, "A")
	n2 := NewNode(context, "B")

	network.AddNode(n1)
	network.AddNode(n2)

	network.AddEdge(n1, n2)

	n1.SetMarg("Wet", 0.64)

	MyMap := make(map[string]string)
	MyMap["A"] = "Wet"

	n2.SetCond("Flood", MyMap, 0.32)
	fmt.Println("A probability is", n1.marg.space["A=Wet"])
	fmt.Println("B probability is", n2.cond.space["B=Flood | A=Wet"])
	fmt.Println("Joint probability is", n1.marg.space["A=Wet"]*n2.cond.space["B=Flood | A=Wet"])
}
