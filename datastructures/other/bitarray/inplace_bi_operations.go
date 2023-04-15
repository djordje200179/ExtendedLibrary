package bitarray

type bitwiseOperation func(a, b uint8) uint8

func (array *Array) applyBiOperation(other *Array, operation bitwiseOperation) {
	if array.Size() != other.Size() {
		panic("Array sizes don't match")
	}

	sliceSize := array.array.Size()
	for i := 0; i < sliceSize; i++ {
		elem1 := array.array.Get(i)
		elem2 := other.array.Get(i)

		resElem := operation(elem1, elem2)
		array.array.Set(i, resElem)
	}
}

func (array *Array) And(other *Array) {
	array.applyBiOperation(other, func(a, b uint8) uint8 {
		return a & b
	})
}

func (array *Array) Or(other *Array) {
	array.applyBiOperation(other, func(a, b uint8) uint8 {
		return a | b
	})
}

func (array *Array) Xor(other *Array) {
	array.applyBiOperation(other, func(a, b uint8) uint8 {
		return a ^ b
	})
}
