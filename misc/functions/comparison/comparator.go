package comparison

type Comparator[T any] func(first, second T) Result
