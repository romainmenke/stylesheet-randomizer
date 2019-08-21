package unit

import mr "math/rand"

func NewSet(options ...func() Value) Value {
	return options[mr.Intn(len(options))]()
}
