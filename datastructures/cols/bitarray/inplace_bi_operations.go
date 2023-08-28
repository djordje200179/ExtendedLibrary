package bitarray

type bitwiseOperation func(a, b uint8) uint8

func (array *Array) applyBiOperation(other *Array, operation bitwiseOperation) {
	if array.Size() != other.Size() {
		panic("Array sizes don't match")
	}

	for i := 0; i < len(array.slice); i++ {
		array.slice[i] = operation(array.slice[i], other.slice[i])
	}
}

// And performs an in-place bitwise AND operation with the other array.
func (array *Array) And(other *Array) {
	array.applyBiOperation(other, func(a, b uint8) uint8 {
		return a & b
	})
}

// Or performs an in-place bitwise OR operation with the other array.
func (array *Array) Or(other *Array) {
	array.applyBiOperation(other, func(a, b uint8) uint8 {
		return a | b
	})
}

// Xor performs an in-place bitwise XOR operation with the other array.
func (array *Array) Xor(other *Array) {
	array.applyBiOperation(other, func(a, b uint8) uint8 {
		return a ^ b
	})
}
