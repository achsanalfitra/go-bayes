package egn

import (
	"fmt"
	"testing"
)

func TestSetJointProbability(t *testing.T) {
	context := BuildContext()

	// Create nodes
	nodeA, err := NewNode(context, "A")
	if err != nil {
		t.Fatalf("Failed to create node A: %v", err)
	}

	nodeB, err := NewNode(context, "B")
	if err != nil {
		t.Fatalf("Failed to create node B: %v", err)
	}

	// Set marginal probability for nodeA: P(A=a) = 0.6
	nodeA.marg.AddPair("A=a", 0.6)
	nodeA.marg.AddPair("A=_", 0.4)
	nodeA.CompleteMarg()
	nodeA.NormalizeMarg()

	// Set conditional probability for nodeB given A: P(B=b|A=a) = 0.7
	parentState := map[string]string{"A": "a"}
	encodedParent := nodeB.encodeParents(parentState)
	if nodeB.cond[encodedParent] == nil {
		nodeB.cond[encodedParent] = NewProbabilitySpace()
	}
	nodeB.cond[encodedParent].AddPair("B=b", 0.7)
	nodeB.cond[encodedParent].AddPair("B=_", 0.3)
	nodeB.CompleteCond()
	nodeB.NormalizeCond()

	// Now set joint probability P(A=a,B=b)
	jointKey := nodeB.encodeFactors(map[string]string{"A": "a", "B": "b"})
	if nodeB.joint[jointKey] == nil {
		nodeB.joint[jointKey] = NewProbabilitySpace()
	}
	nodeB.joint[jointKey].AddPair("joint_event", 0.42) // For example 0.6 * 0.7 = 0.42 joint prob

	// Check joint probability
	jointProb := nodeB.joint[jointKey].TotalProb()
	if jointProb != 0.42 {
		t.Errorf("Expected joint probability 0.42, got %f", jointProb)
	}

	// Optional: print to confirm
	fmt.Printf("Joint probability set: %v\n", jointProb)
}
