package misc

import "fmt"

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func (pair Pair[T1, T2]) Get() (T1, T2) { return pair.First, pair.Second }

func (pair Pair[T1, T2]) String() string {
	return fmt.Sprintf("(%v, %v)", pair.First, pair.Second)
}
