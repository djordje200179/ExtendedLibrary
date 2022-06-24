package comparison

func Ascending[T ordered](a, b T) Result {
	switch {
	case a < b:
		return FirstSmaller
	case a > b:
		return SecondSmaller
	default:
		return Equal
	}
}

func Descending[T ordered](a, b T) Result {
	switch {
	case a < b:
		return FirstBigger
	case a > b:
		return SecondBigger
	default:
		return Equal
	}
}
