package hlp

import "fmt"

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

func (b *BiMapInt) AddKey(key string) error {
	// safeguard against overwrite
	if _, exist := b.StrInt[key]; exist {
		return fmt.Errorf("key already exists")
	}

	// one source of truth for int id
	id := len(b.StrInt)
	b.StrInt[key] = id
	b.IntStr[id] = key

	return nil
}
