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

func (b *BiMapInt) DeleteKey(key string) error {
	// safeguard against missing key
	if _, exist := b.StrInt[key]; !exist {
		return fmt.Errorf("key %s doesn't exist", key)
	}

	id := b.StrInt[key]

	// delete the key from both maps
	delete(b.StrInt, key)
	delete(b.IntStr, id)

	// reindex the StrInt map
	for str, idx := range b.StrInt {
		if idx > id {
			b.StrInt[str] = idx - 1
		}
	}

	// recreate the IntStr map
	b.IntStr = make(map[int]string, len(b.StrInt))
	for str, idx := range b.StrInt {
		b.IntStr[idx] = str
	}

	return nil
}
