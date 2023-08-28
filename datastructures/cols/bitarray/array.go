package bitarray

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"strings"
)

// Array is an optimized array of boolean values.
// The zero value is ready to use. Do not copy a non-zero Array.
type Array struct {
	slice       []uint8
	lastElemOff uint8
}

// New creates an empty Array.
func New() *Array {
	return &Array{
		slice: make([]uint8, 0),
	}
}

// NewWithSize creates an empty Array with the specified initial size.
func NewWithSize(initialSize int) *Array {
	arrSize := (initialSize + 7) / 8

	return &Array{
		slice:       make([]uint8, arrSize),
		lastElemOff: uint8(initialSize % 8),
	}
}

// NewWithCapacity creates an empty Array with the specified initial capacity.
func NewWithCapacity(initialCapacity int) *Array {
	arrCapacity := (initialCapacity + 7) / 8

	return &Array{
		slice:       make([]uint8, 0, arrCapacity),
		lastElemOff: 0,
	}
}

// NewFromSlice creates an Array from the specified slice.
func NewFromSlice(slice []bool) *Array {
	arr := NewWithSize(len(slice))

	for i := 0; i < len(slice); i++ {
		arr.Set(i, slice[i])
	}

	return arr
}

// Size returns the number of bits in the Array.
func (array *Array) Size() int {
	if array.lastElemOff == 0 {
		return len(array.slice) * 8
	} else {
		return (len(array.slice)-1)*8 + int(array.lastElemOff)
	}
}

// Capacity returns the number of bits that the Array can hold without reallocating.
func (array *Array) Capacity() int {
	return cap(array.slice) * 8
}

func (array *Array) getRealIndex(index int) int {
	size := array.Size()

	if index >= size || index < -size {
		cols.PanicOnIndexOutOfBounds(index, size)
	}

	if index < 0 {
		index += size
	}

	return index
}

// Get returns the bit at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (array *Array) Get(index int) bool {
	index = array.getRealIndex(index)

	elemIndex := index / 8
	elemOff := index % 8
	masked := array.slice[elemIndex] & (1 << elemOff)

	return masked != 0
}

// Set sets the bit at the specified index to the specified value.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (array *Array) Set(index int, value bool) {
	index = array.getRealIndex(index)

	elemIndex := index / 8
	elemOff := index % 8
	elem := array.slice[elemIndex]

	mask := uint8(1) << elemOff
	if value {
		elem |= mask
	} else {
		elem &= ^mask
	}

	array.slice[elemIndex] = elem
}

// SetAll sets all bits in the Array to the specified value.
func (array *Array) SetAll(value bool) {
	var elem uint8
	if value {
		elem = 0xFF
	}

	for i := 0; i < len(array.slice); i++ {
		array.slice[i] = elem
	}
}

// Flip flips the bit at the specified index.
func (array *Array) Flip(index int) {
	index = array.getRealIndex(index)

	elemIndex := index / 8
	elemOff := index % 8
	elem := array.slice[elemIndex]

	mask := uint8(1) << elemOff
	elem ^= mask

	array.slice[elemIndex] = elem
}

// FlipAll flips all bits in the Array.
func (array *Array) FlipAll() {
	for i := 0; i < len(array.slice); i++ {
		array.slice[i] = ^array.slice[i]
	}
}

// Append appends the specified bit value to the end of the Array.
func (array *Array) Append(value bool) {
	if array.lastElemOff == 0 {
		array.slice = append(array.slice, 0)
	}

	array.lastElemOff = (array.lastElemOff + 1) % 8

	array.Set(array.Size()-1, value)
}

// Insert inserts the specified bit value at the specified index.
func (array *Array) Insert(index int, value bool) {
	index = array.getRealIndex(index)

	elemIndex := index / 8
	elemOff := index % 8

	if array.lastElemOff == 0 {
		array.slice = append(array.slice, 0)
	}
	array.lastElemOff = (array.lastElemOff + 1) % 8

	var lastBit bool
	if elemOff != 0 {
		elem := array.slice[elemIndex]

		lastBit = elem&(1<<7) != 0
		higherBits := elem >> elemOff << (elemOff + 1)
		lowerBits := elem << (8 - elemOff) >> (8 - elemOff)

		var shiftedValue uint8
		if value {
			shiftedValue = 1 << elemOff
		}

		newElem := higherBits | lowerBits | shiftedValue

		array.slice[elemIndex] = newElem
	}

	for i := elemIndex + 1; i < len(array.slice); i++ {
		elem := array.slice[i]

		newLastBit := elem&(1<<7) != 0
		elem = elem << 1

		if lastBit {
			elem |= 1
		}

		array.slice[i] = elem
		lastBit = newLastBit
	}
}

// Remove removes the bit at the specified index.
func (array *Array) Remove(index int) {
	index = array.getRealIndex(index)

	elemIndex := index / 8
	elemOff := index % 8

	var firstBit bool
	for i := len(array.slice) - 1; i > elemIndex; i-- {
		elem := array.slice[i]

		newFirstBit := elem&1 != 0
		elem = elem >> 1

		if firstBit {
			elem |= 1 << 7
		}

		array.slice[i] = elem
		firstBit = newFirstBit
	}

	if elemOff != 0 {
		elem := array.slice[elemIndex]

		higherBits := elem >> elemOff << elemOff
		lowerBits := elem << (8 - elemOff) >> (8 - elemOff)

		newElem := higherBits | lowerBits

		array.slice[elemIndex] = newElem
	}

	array.lastElemOff = (array.lastElemOff - 1) % 8
	if array.lastElemOff == 0 {
		array.slice = array.slice[:len(array.slice)-1]
	}
}

// Clear removes all bits from the Array.
func (array *Array) Clear() {
	array.slice = make([]uint8, 0)
	array.lastElemOff = 0
}

// Reverse reverses the order of the bits in the Array.
func (array *Array) Reverse() {
	panic("not implemented")
}

// Join appends all bits from the specified Array to the end of the Array.
func (array *Array) Join(other *Array) {
	if array.lastElemOff == 0 {
		array.slice = append(array.slice, other.slice...)
		array.lastElemOff = other.lastElemOff
	} else {
		for i := 0; i < other.Size(); i++ {
			array.Append(other.Get(i))
		}
	}

	other.Clear()
}

// Clone returns a shallow copy of the Array.
func (array *Array) Clone() *Array {
	cloned := NewWithSize(array.Size())
	cloned.lastElemOff = array.lastElemOff
	copy(cloned.slice, array.slice)

	return cloned
}

// String returns a string representation of the Array.
func (array *Array) String() string {
	var sb strings.Builder

	for i := 0; i < array.Size(); i++ {
		val := array.Get(i)

		var char byte
		if val {
			char = '1'
		} else {
			char = '0'
		}

		sb.WriteByte(char)
	}

	return sb.String()
}
