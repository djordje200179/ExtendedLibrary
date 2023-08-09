# Streams

This package lets you use Java-like generic streams in Go.
Streams can be either infinite or finite.
Values are lazy-fetched, and therefore, values aren't processed until
terminal methods are invoked.

## Construction
You can create a stream in multiple ways:
1. Manually (by creating and setting supplier)
2. Streaming object (by calling `Stream` method on an object that implements `Streamer` interface)
3. Constructor (by calling one of the provided constructor functions):
   - `FromChannel(channel <-chan T) Stream[T]`
   - `FromValues(values ...T) Stream[T]`
   - `FromRange(lower, upper int) Stream[int]`
   - `FromSlice(slice []T) Stream[T]`
   - `FromSliceRefs(slice []T) Stream[*T]`
   - `FromMap(m map[K]V) Stream[misc.Pair[K, V]]`
   - `FromFiniteGenerator(generator functions.EmptyGenerator[optional.Optional[T]]) Stream[T]`
   - `FromInfiniteGenerator(generator functions.EmptyGenerator[T]) Stream[T]`

## Intermediates
There are methods that you can use to transform the stream:
- `Map(mapper functions.Mapper[T, U]) Stream[U]`
- `Filter(predicate functions.predicate[T]) Stream[T]`
- `Limit(count int) Stream[T]`
- `Seek(count int) Stream[T]`
- `Sort(comparator functions.Comparator[T]) Stream[T]`
- `Window(width int) Stream[[]T]`

## Terminals
After you have transformed the stream, you can use terminal methods
to process values in the stream and get the result that you wanted:
- `ForEach(function functions.ParamCallback[T])`
- `Reduce(accumulator P, reducer functions.Reducer[T, P]) P`
- `Any(predicate functions.predicate[T]) bool`
- `All(predicate functions.predicate[T]) bool`
- `Collect[T, R any](collector Collector[T, R]) R`
- `Count() int`
- `Max(comparator functions.Comparator[T]) optional.Optional[T]`
- `Min(comparator functions.Comparator[T]) optional.Optional[T]`
- `First() optional.Optional[T]`
- `Find(predicate functions.predicate[T]) optional.Optional[T]`

Because streams are very similar to built-in channels, there is
a method for converting the stream into a channel:
`Channel() <-chan T`

## Limitations
Due to lack of generic methods in Go, some methods (like `Map` and `Reduce`) 
currently only support returning same type like stream values are.
If your transformations return some other type, then you should consider 
functions with the same name that accept the stream as the first argument.

```go
//Instead of
stream.Map(mapFunc)

//You should use
streams.Map(stream, mapFunc)
```

Also, because of cyclic generic instantiations, some methods (like `Window`)
are also functions that accept the stream as first argument.