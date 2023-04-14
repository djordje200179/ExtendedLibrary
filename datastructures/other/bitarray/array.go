package bitarray

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/collections/array"
)

type Array struct {
	array       array.Array[uint8]
	lastElemOff uint8
}

func New() *Array {
	return &Array{
		array:       *array.New[uint8](),
		lastElemOff: 0,
	}
}

func NewWithSize(initialSize int) *Array {
	arrSize := (initialSize + 7) / 8

	return &Array{
		array:       *array.NewWithSize[uint8](arrSize),
		lastElemOff: uint8(initialSize % 8),
	}
}

func NewWithCapacity(initialCapacity int) *Array {
	arrCapacity := (initialCapacity + 7) / 8

	return &Array{
		array:       *array.NewWithCapacity[uint8](arrCapacity),
		lastElemOff: 0,
	}
}

func FromSlice(slice []bool) *Array {
	arr := NewWithSize(len(slice))

	for i := 0; i < len(slice); i++ {
		arr.Set(i, slice[i])
	}

	return arr
}

func (array *Array) Size() int {
	if array.lastElemOff == 0 {
		return array.array.Size() * 8
	} else {
		return (array.array.Size()-1)*8 + int(array.lastElemOff)
	}
}

func (array *Array) getRealIndex(index int) int {
	size := array.Size()

	if index >= size || index < -size {
		panic(fmt.Sprintf("Index out of bounds: %d", index))
	}

	if index < 0 {
		index += size
	}

	return index
}

func (array *Array) Get(index int) bool {
	index = array.getRealIndex(index)

	elemIndex := index / 8
	elemOff := index % 8
	elem := array.array.Get(elemIndex)
	masked := elem & (1 << elemOff)

	return masked != 0
}

func (array *Array) Set(index int, value bool) {
	index = array.getRealIndex(index)

	elemIndex := index / 8
	elemOff := index % 8
	elem := array.array.Get(elemIndex)

	mask := uint8(1) << elemOff
	if value {
		elem |= mask
	} else {
		elem &= ^mask
	}

	array.array.Set(elemIndex, elem)
}

func (array *Array) Append(value bool) {
	if array.lastElemOff == 0 {
		array.array.Append(0)
	}

	array.lastElemOff = (array.lastElemOff + 1) % 8

	array.Set(array.Size()-1, value)
}

func (array *Array) Insert(index int, value bool) {
	index = array.getRealIndex(index)

	elemIndex := index / 8
	elemOff := index % 8

	if array.lastElemOff == 0 {
		array.array.Append(0)
	}
	array.lastElemOff = (array.lastElemOff + 1) % 8

	var lastBit bool
	if elemOff != 0 {
		elem := array.array.Get(elemIndex)

		lastBit = elem&(1<<7) != 0
		higherBits := elem >> elemOff << (elemOff + 1)
		lowerBits := elem << (8 - elemOff) >> (8 - elemOff)

		var shiftedValue uint8
		if value {
			shiftedValue = 1 << elemOff
		}

		newElem := higherBits | lowerBits | shiftedValue

		array.array.Set(elemIndex, newElem)
	}

	for i := elemIndex + 1; i < array.array.Size(); i++ {
		elem := array.array.Get(i)

		newLastBit := elem&(1<<7) != 0
		elem = elem << 1

		if lastBit {
			elem |= 1
		}

		array.array.Set(i, elem)
		lastBit = newLastBit
	}
}

func (array *Array) Remove(index int) {
	index = array.getRealIndex(index)

	elemIndex := index / 8
	elemOff := index % 8

	var firstBit bool
	for i := array.array.Size() - 1; i > elemIndex; i-- {
		elem := array.array.Get(i)

		newFirstBit := elem&1 != 0
		elem = elem >> 1

		if firstBit {
			elem |= 1 << 7
		}

		array.array.Set(i, elem)
		firstBit = newFirstBit
	}

	if elemOff != 0 {
		elem := array.array.Get(elemIndex)

		higherBits := elem >> elemOff << elemOff
		lowerBits := elem << (8 - elemOff) >> (8 - elemOff)

		newElem := higherBits | lowerBits

		array.array.Set(elemIndex, newElem)
	}

	array.lastElemOff = (array.lastElemOff - 1) % 8
	if array.lastElemOff == 0 {
		array.array.Remove(array.array.Size() - 1)
	}
}

func (array *Array) Clear() {
	array.array.Clear()
	array.lastElemOff = 0
}

func (array *Array) Reverse() {
	panic("Not implemented")
}

func (array *Array) Swap(index1, index2 int) {
	elem1 := array.Get(index1)
	elem2 := array.Get(index2)

	array.Set(index1, elem2)
	array.Set(index2, elem1)
}

func (array *Array) Join(other *Array) {
	if array.lastElemOff == 0 {
		array.array.AppendMany(other.array...)
		array.lastElemOff = other.lastElemOff
	} else {
		for i := 0; i < other.Size(); i++ {
			array.Append(other.Get(i))
		}
	}

	other.Clear()
}

func (array *Array) Clone() *Array {
	cloned := NewWithSize(array.Size())
	cloned.lastElemOff = array.lastElemOff
	copy(cloned.array, array.array)

	return cloned
}
