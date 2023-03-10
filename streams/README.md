# streams

This package lets you use Java-like generic streams in Go.
Streams can be either infinite or finite, and values are 
lazy-fetched. Therefore, values aren't processed until
terminal methods are invoked.

You can create a stream manually by setting supplier or using 
many constructor functions:
- `FromChannel(channel <-chan T) Stream[T]`
- `FromValues(values ...T) Stream[T]`
- `FromRange(lower, upper int) Stream[int]`
- `FromSlice(slice []T) Stream[T]`
- `FromSliceRefs(slice []T) Stream[*T]`
- `FromMap(m map[K]V) Stream[misc.Pair[K, V]]`
- `FromFiniteGenerator(generator functions.EmptyGenerator[optional.Optional[T]]) Stream[T]`
- `FromInfiniteGenerator(generator functions.EmptyGenerator[T]) Stream[T]`

There are methods that you can use to transform the stream:
- `Map(mapper functions.Mapper[T, U]) Stream[U]`
- `Filter(predictor functions.Predictor[T]) Stream[T]`
- `Limit(count int) Stream[T]`
- `Seek(count int) Stream[T]`
- `Sort(comparator functions.Comparator[T]) Stream[T]`

At the end, you use terminal methods to process values in
the stream and get the result that you wanted:
- `ForEach(function functions.ParamCallback[T])`
- `Reduce(accumulator P, reducer functions.Reducer[T, P]) P`
- `Any(predictor functions.Predictor[T]) bool`
- `All(predictor functions.Predictor[T]) bool`
- `Collect[T, R any](collector Collector[T, R]) R`
- `Count() int`
- `Max(comparator functions.Comparator[T]) optional.Optional[T]`
- `Min(comparator functions.Comparator[T]) optional.Optional[T]`
- `First() optional.Optional[T]`
- `Find(predictor functions.Predictor[T]) optional.Optional[T]`

Because streams are very similar to built-in channels, there is
a method for converting the stream into a channel:
`Channel() <-chan T`


**NOTE:**
Due to limitations of Go regarding generic method, methods 
that should be generic (like map and reduce) are functions that
accept the stream as first argument. 

Hopefully in future versions of Go there will be generic methods,
and this package will be updated to fix these issues.

For example, instead of `stream.Map(mapFunc)` you should be using 
`streams.Map(stream, mapFunc)`.
