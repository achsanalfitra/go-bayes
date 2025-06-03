package hlp

type BiMapInt struct {
	StrInt map[string]int
	IntStr map[int]string
}

func NewBiMapInt() *BiMapInt {
	return &BiMapInt{
		StrInt: make(map[string]int),
		IntStr: make(map[int]string),
	}
}
