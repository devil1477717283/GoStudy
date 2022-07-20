package generic

type Addable interface {
	string | int | float64 | float32
}

func Add[T Addable](a, b T) T {
	return a + b
}
