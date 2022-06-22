package bitarray

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
)

type BitArray struct {
	array       array.Array[uint8]
	lastElemOff uint8
}

func New() *BitArray {
	return &BitArray{
		array:       *array.New[uint8](),
		lastElemOff: 0,
	}
}

func (b *BitArray) Size() int {
	return (b.array.Size()-1)*8 + int(b.lastElemOff)
}

func (b *BitArray) Get(index int) bool {
	if index >= b.Size() {
		panic(fmt.Sprintf("runtime error: index out of range [%d] with length %d", index, b.Size()))
	}

	elemIndex := index / 8
	elemOff := index % 8
	elem := b.array.Get(elemIndex)
	masked := elem & (1 << elemOff)

	return masked != 0
}

func (b *BitArray) Set(index int, value bool) {
	if index >= b.Size() {
		panic(fmt.Sprintf("runtime error: index out of range [%d] with length %d", index, b.Size()))
	}

	elemIndex := index / 8
	elemOff := index % 8
	elem := b.array.Get(elemIndex)

	mask := uint8(1) << elemOff
	if value {
		elem |= mask
	} else {
		elem &= ^mask
	}

	b.array.Set(elemIndex, elem)
}

func (b *BitArray) Append(value bool) {
	if b.lastElemOff == 0 {
		b.array.Append(0)
	}

	b.lastElemOff = (b.lastElemOff + 1) % 8

	b.Set(b.Size()-1, value)
}

func (b *BitArray) Insert(index int, value bool) {
	elemIndex := index / 8
	elemOff := index % 8
	elem := b.array.Get(elemIndex)

}

func (b *BitArray) Remove(index int) bool {
	elemIndex := index / 8
	elemOff := index % 8
	elem := b.array.Get(elemIndex)
}
