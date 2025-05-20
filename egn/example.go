package egn

import "fmt"

func Example() {
	context := BuildContext()

	nodeA, err := NewNode(context, "A")
	if err != nil {
		fmt.Println(err)
	}

	nodeA.Set.MarginalProbability("a1", 0.5)
	nodeA.Set.MarginalProbability("a2", 0.3)
	nodeA.Set.MarginalProbability("a3", 0.2)
	nodeA.Set.MarginalProbability("a3", 0.2)

	nodeA.Show.MarginalEvents()

	nodeB, err := NewNode(context, "A")
	if err != nil {
		fmt.Println(err)
	}

	nodeB.Set.MarginalProbability("b1", 0.5)
	nodeB.Set.MarginalProbability("b2", 0.3)
	nodeB.Set.MarginalProbability("b3", 0.2)

	nodeB.Show.MarginalEvents()
}
