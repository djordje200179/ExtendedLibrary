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

type CopyOnWriteArray[T any] struct {
	rawArray *rawArray.Array[T]
	mutex    sync.Mutex
}

func New[T any]() *CopyOnWriteArray[T] {
	return NewWithSize[T](0)
}

func NewWithCapacity[T any](capacity int) *CopyOnWriteArray[T] {
	return FromArray(rawArray.NewWithCapacity[T](capacity))
}

func NewWithSize[T any](size int) *CopyOnWriteArray[T] {
	return FromArray(rawArray.NewWithSize[T](size))
}

func FromArray[T any](array *rawArray.Array[T]) *CopyOnWriteArray[T] {
	return &CopyOnWriteArray[T]{rawArray: array}
}

func FromSlice[T any](slice []T) *CopyOnWriteArray[T] {
	return FromArray(rawArray.FromSlice(slice))
}

func Collector[T any]() streams.Collector[T, collections.Collection[T]] {
	return collections.Collector[T]{New[T]()}
}

func (array *CopyOnWriteArray[T]) Size() int {
	return array.rawArray.Size()
}

func (array *CopyOnWriteArray[T]) Get(index int) T {
	return array.rawArray.Get(index)
}

func (array *CopyOnWriteArray[T]) GetRef(index int) *T {
	return array.rawArray.GetRef(index)
}

func (array *CopyOnWriteArray[T]) Set(index int, value T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Set(index, value)

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) Append(value T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Append(value)

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) AppendMany(values ...T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.AppendMany(values...)

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) Insert(index int, value T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Insert(index, value)

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) InsertMany(index int, values ...T) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.InsertMany(index, values...)

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) Remove(index int) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Remove(index)

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) Clear() {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Clear()

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) Reverse() {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Reverse()

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) Sort(comparator comparison.Comparator[T]) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Sort(comparator)

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) Join(other collections.Collection[T]) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	newArray := array.rawArray.Clone().(*rawArray.Array[T])
	newArray.Join(other)

	array.rawArray = newArray
}

func (array *CopyOnWriteArray[T]) Clone() collections.Collection[T] {
	clonedRawArray := array.rawArray.Clone().(*rawArray.Array[T])

	return &CopyOnWriteArray[T]{rawArray: clonedRawArray}
}

func (array *CopyOnWriteArray[T]) Iterator() iterable.Iterator[T] {
	return array.ModifyingIterator()
}

func (array *CopyOnWriteArray[T]) ModifyingIterator() collections.Iterator[T] {
	return &Iterator[T]{array, 0}
}

func (array *CopyOnWriteArray[T]) Stream() streams.Stream[T] {
	return array.rawArray.Stream()
}

func (array *CopyOnWriteArray[T]) RefsStream() streams.Stream[*T] {
	return array.rawArray.RefsStream()
}

func (array *CopyOnWriteArray[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	return array.rawArray.FindIndex(predicate)
}

func (array *CopyOnWriteArray[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	return array.rawArray.FindRef(predicate)
}

func (array *CopyOnWriteArray[T]) Slice() []T {
	return array.rawArray.Slice()
}

func (array *CopyOnWriteArray[T]) SliceRange(from, to int) []T {
	return array.rawArray.SliceRange(from, to)
}
