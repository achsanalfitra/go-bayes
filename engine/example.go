package engine

import "fmt"

func Test() {
	context := BuildContext()

	n1 := NewNode(context, "A")
	n2 := NewNode(context, "B")

	n2.AddParent(n1)

	n1.SetMarg("Wet", 0.64)

	MyMap := make(map[string]string)
	MyMap["A"] = "Wet"

	n2.SetCond("Flood", MyMap, 0.32)
	fmt.Println("A probability is", n1.marg.space["A=Wet"])
	fmt.Println("B probability is", n2.cond.space["B=Flood | A=Wet"])
	fmt.Println("Joint probability is", n1.marg.space["A=Wet"]*n2.cond.space["B=Flood | A=Wet"])
}
