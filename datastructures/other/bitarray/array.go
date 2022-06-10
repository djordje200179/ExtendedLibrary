package bitarray

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
	"github.com/djordje200179/extendedlibrary/streams"
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
	elementIndex := index / 8
	elementOffset := index % 8
	element := b.array.Get(elementIndex)

	return element & (1 << elementOffset)
}

func (b *BitArray) Set(index int, value bool) {
	elementIndex := index / 8
	elementOffset := index % 8
	element := b.array.Get(elementIndex)

	if value {
		element |= 1 << elementOffset
	} else {
		element &= ^(1 << elementOffset)
	}

	b.array.Set(elementIndex, element)
}

func (b *BitArray) Append(value bool) {
	elementIndex := b.array.Size() - 1
	elementOffset := b.lastElemOff

}

func (b *BitArray) Insert(index int, value bool) {

}

func (b *BitArray) Remove(index int) bool {

}

func (b *BitArray) Stream() streams.Stream[bool] {

}
