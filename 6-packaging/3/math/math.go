package math

type Math struct {
	A int
	B int
}

func (math *Math) Add() int {
	return math.A + math.B
}
