package bitarray

import "errors"

type bitwiseOp func(a, b uint8) uint8

var SizeMismatchError = errors.New("array sizes don't match")

func (array *Array) applyBiOperation(other *Array, operation bitwiseOp) {
	if array.Size() != other.Size() {
		panic(SizeMismatchError)
	}

	for i, val := range array.slice {
		array.slice[i] = operation(val, other.slice[i])
	}
}

// And performs an in-place bitwise AND operation with the other Array.
//
// If the arrays are not of the same size, panic SizeMismatchError occurs.
func (array *Array) And(other *Array) {
	array.applyBiOperation(other, func(a, b uint8) uint8 {
		return a & b
	})
}

// Or performs an in-place bitwise OR operation with the other Array.
//
// If the arrays are not of the same size, panic SizeMismatchError occurs.
func (array *Array) Or(other *Array) {
	array.applyBiOperation(other, func(a, b uint8) uint8 {
		return a | b
	})
}

// Xor performs an in-place bitwise XOR operation with the other Array.
//
// If the arrays are not of the same size, panic SizeMismatchError occurs.
func (array *Array) Xor(other *Array) {
	array.applyBiOperation(other, func(a, b uint8) uint8 {
		return a ^ b
	})
}
