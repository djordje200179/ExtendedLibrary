package functions

import (
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
)

type ParamCallback[T any] func(value T)
type EmptyCallback func()

type ParamGenerator[T, P any] func(value P) T
type EmptyGenerator[T any] func() T

type Reducer[T, P any] func(acc P, value T) P
type Mapper[T, P any] func(value T) P

type Comparator[T any] func(first, second T) comparison.Result
