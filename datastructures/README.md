# Data structures

Java-like object-oriented data structures for Go.
Every type implements the most commonly used methods for it.

## Types
This package implements many generic data structures:
- Collections
	- Array
	- Linked list
	- Bitarray
    - Concurrent linked list
	- Read-only wrapper
	- Concurrent wrapper
- Maps
	- Red-black tree (binary search tree)
	- Hashmap
	- Linked wrapper
	- Read-only wrapper
	- Concurrent wrapper
- Sets
	- Hashset
    - Tree set
    - Bitarray set
	- Read-only wrapper
- Sequences
	- Bounded buffer
    - Linked list deque
    - Array deque
	- Priority queue
- Other
	- Matrix

## Iteration

### Iterators
Collection, maps and sets support iterations through common `Iterable` interface,
that fits well into `for` loop.

```go
for it := list.Iterator(); it.Valid(); it.Next() {
	value := it.Get();
}
```
Be aware that this kind of iterator supports only reading elements, not modifying them.
To modify elements, you need to get a special iterator for that kind of collection.
They also have specialized methods to modify structure (like insertion before the
current element in a linked list).
And through them, you can also access the element directly by reference (pointer).

### Streams
Collection, maps and sets support value streaming through Go 1.22 
range over func functionality.

```go
for val := range collection.Stream {
	
}
```

Collections and maps also support getting a stream of references (pointers) 
to elements, so you don't need to copy huge structure elements

## Construction

### Constructors
Every type has one or more constructors that allow you to create new instances of
the type.

```go
arr1 := array.New[T]()
arr2 := array.NewWithCapacity[T](capacity)
hmap := hashmap.NewWithCapacity[T](capacity)
list := linklist.NewFromIterable[T](iterable)
```

### Casts
Some types use others as their internal representation (slices, built-in maps, etc.).
Therefore, those types have constructors that allow you to create them from underlying types.

```go
arr := array.FromSlice[T](slice)
set := mapset.FromMap(hashmap)
queue := boundedbuffer.FromChannel(ch)
```

### Streams
You can also collect every finite stream into a suitable collection.

```go
stream := streams.New(suppliers.Range(0, 100))
iter := iter.StreamIterable(stream)
arr := array.NewFromIterable[T](iter)
```