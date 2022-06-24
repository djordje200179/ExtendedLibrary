package misc

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func (p Pair[T1, T2]) Get() (T1, T2) { return p.First, p.Second }
