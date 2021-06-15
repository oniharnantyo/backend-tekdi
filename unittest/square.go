package unittest

type square struct {
	Length int
}

func NewSquare(length int) *square {
	return &square{
		Length: length,
	}
}

func (s *square) Area() int {
	return s.Length * s.Length
}

func (s *square) Around() int {
	return 4 * s.Length
}
