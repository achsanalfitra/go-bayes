package egn

type Validator interface {
	CheckNode()         //	individual check
	CheckMarginal()     // check marginal consistency
	CheckConditional()  // check conditional consistency
	CheckJoint()        // check the existence of complete joint
	CheckCompleteness() // check the completeness of all
	CheckInferrable()   // check whether the context is ready for inference or not
}
