package copyonwritearray

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	rawArray "github.com/djordje200179/extendedlibrary/datastructures/collections/array"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Array[T any] struct {
	rawArray *rawArray.Array[T]
	mutex    sync.Mutex
}

func New[T any]() *Array[T] {
	return NewWithSize[T](0)
}

func NewWithCapacity[T any](capacity int) *Array[T] {
	return FromArray(rawArray.NewWithCapacity[T](capacity))
}

func NewWithSize[T any](size int) *Array[T] {
	return FromArray(rawArray.NewWithSize[T](size))
}

func FromArray[T any](array *rawArray.Array[T]) *Array[T] {
	return &Array[T]{rawArray: array}
}

func FromSlice[T any](slice []T) *Array[T] {
	return FromArray(rawArray.FromSlice(slice))
}

func Collector[T any]() streams.Collector[T, *Array[T]] {
	return collections.Collector[T, *Array[T]]{New[T]()}
}

func (array *Array[T]) Size() int {
	return array.rawArray.Size()
}

func (array *Array[T]) Get(index int) T {
	return array.rawArray.Get(index)
}

func (array *Array[T]) GetRef(index int) *T {
	return array.rawArray.GetRef(index)
}

func (array *Array[T]) Set(index int, value T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Set(index, value)

	array.rawArray = newArray
}

func (array *Array[T]) Append(value T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Append(value)

	array.rawArray = newArray
}

func (array *Array[T]) AppendMany(values ...T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.AppendMany(values...)

	array.rawArray = newArray
}

func (array *Array[T]) Insert(index int, value T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Insert(index, value)

	array.rawArray = newArray
}

func (array *Array[T]) InsertMany(index int, values ...T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.InsertMany(index, values...)

	array.rawArray = newArray
}

func (array *Array[T]) Remove(index int) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Remove(index)

	array.rawArray = newArray
}

func (array *Array[T]) Clear() {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Clear()

	array.rawArray = newArray
}

func (array *Array[T]) Reverse() {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Reverse()

	array.rawArray = newArray
}

func (array *Array[T]) Sort(comparator comparison.Comparator[T]) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Sort(comparator)

	array.rawArray = newArray
}

func (array *Array[T]) Join(other collections.Collection[T]) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Join(other)

	array.rawArray = newArray
}

func (array *Array[T]) Clone() collections.Collection[T] {
	clonedRawArray := array.rawArray.Clone().(*rawArray.Array[T])

	return &Array[T]{rawArray: clonedRawArray}
}

func (array *Array[T]) Iterator() iterable.Iterator[T] {
	return array.CollectionIterator()
}

func (array *Array[T]) CollectionIterator() collections.Iterator[T] {
	return &Iterator[T]{array, 0}
}

func (array *Array[T]) Stream() streams.Stream[T] {
	return array.rawArray.Stream()
}

func (array *Array[T]) RefsStream() streams.Stream[*T] {
	return array.rawArray.RefsStream()
}

func (array *Array[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	return array.rawArray.FindIndex(predicate)
}

func (array *Array[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	return array.rawArray.FindRef(predicate)
}

func (array *Array[T]) Slice() []T {
	return array.rawArray.Slice()
}

func (array *Array[T]) SliceRange(from, to int) []T {
	return array.rawArray.SliceRange(from, to)
}
