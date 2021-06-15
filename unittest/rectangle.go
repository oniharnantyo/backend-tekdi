package unittest

type rectangle struct {
	Length int
	Wide   int
}

func NewRectangle(length, wide int) *rectangle {
	return &rectangle{
		Length: length,
		Wide:   wide,
	}
}

func (r *rectangle) Area() int {

}

func (r *rectangle) Around() int {

}
