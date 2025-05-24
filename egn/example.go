package egn

import (
	"fmt"
	"log"
)

func Example() {
	context := BuildContext()

	nodeA, err := NewNode(context, "A")
	if err != nil {
		log.Println(err)
		return
	}
	nodeB, err := NewNode(context, "B")
	if err != nil {
		fmt.Println(err)
		return
	}

	nodeA.Set.MarginalProbability("a1", 0.5)
	nodeA.Set.MarginalProbability("a2", 0.3)
	nodeA.Set.MarginalProbability("a3", 0.2)

	givenBb1 := map[string]string{
		"B": "b1",
	}

	nodeA.Set.ConditionalProbability("a1", givenBb1, 0.5)
	nodeA.Set.ConditionalProbability("a2", givenBb1, 0.3)
	nodeA.Set.ConditionalProbability("a3", givenBb1, 0.2)

	nodeB.Set.MarginalProbability("b1", 0.5)
	nodeB.Set.MarginalProbability("b2", 0.3)
	nodeB.Set.MarginalProbability("b3", 0.2)

	context.ShowMarginal()
	nodeA.Show.ConditionalEvents()
}
