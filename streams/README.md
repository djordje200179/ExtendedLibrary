# Streams

This package lets you use Java-like generic streams in Go.
Streams can be either infinite or finite.
Values are lazy-fetched, and therefore, values aren't processed 
until terminal methods are invoked.

It is worth to mention that there are two types of streams:
- `Stream[T]` - for single-value streams
- `Stream2[T, U]` - for streams that return two values at once

## Construction
You can create a stream in multiple ways:
1. Manually (by calling `From` or `From2` method with a seq function as argument)
2. Streaming object (by calling `Stream` or `Stream2` method on an object that 
implements `Streamer` or `Streamer2` interface)
3. Constructor (by calling one of the provided constructor functions):
	- `FromChannel(ch <-chan T) Stream[T]`
	- `Range(lower, upper int) Stream[int]`
	- `FromGenerator[T any](generator func() T) Stream[T]`

## Intermediates
There are methods that you can use to transform
single-value streams:
- `Map(mapper functions.Mapper[T, U]) Stream[U]`
- `Filter(predicate functions.predicate[T]) Stream[T]`
- `Limit(count int) Stream[T]`
- `Seek(count int) Stream[T]`
- `Sort(comparator functions.Comparator[T]) Stream[T]`
- `Window(width int) Stream[[]T]`

And methods that you can use to transform two-value streams:
- `Keys() Stream[K]`
- `Values() Stream[V]`
- `FilterKeys(predicate functions.predicate[K]) Stream2[K, V]`
- `FilterValues(predicate functions.predicate[V]) Stream2[K, V]`
- `Limit(count int) Stream2[K, V]`
- `Seek(count int) Stream2[K, V]`

## Terminals
After you have transformed the stream, you can use terminal methods
to process values in the stream and get the result that you wanted.

Methods that you can use to process single-value streams:
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
- `Channel() <-chan T`

Methods that you can use to process two-value streams:
- `Count() int`

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