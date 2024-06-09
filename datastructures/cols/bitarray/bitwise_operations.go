package bitarray

// And performs bitwise AND operation and returns resulting Array.
func And(array1, array2 *Array) *Array {
	resArray := array1.Clone()
	resArray.And(array2)
	return resArray
}

// Or performs bitwise OR operation and returns resulting Array.
func Or(array1, array2 *Array) *Array {
	resArray := array1.Clone()
	resArray.Or(array2)
	return resArray
}

// Xor performs bitwise XOR operation and returns resulting Array.
func Xor(array1, array2 *Array) *Array {
	resArray := array1.Clone()
	resArray.Xor(array2)
	return resArray
}

// Not performs bitwise NOT operation and returns resulting Array.
func Not(array *Array) *Array {
	result := array.Clone()
	result.FlipAll()
	return result
}

// ShiftLeft performs bitwise left shift operation and returns resulting Array.
func ShiftLeft(array *Array, operationType ShiftType) *Array {
	result := array.Clone()
	result.ShiftLeft(operationType)
	return result
}

// ShiftRight performs bitwise right shift operation and returns resulting Array.
func ShiftRight(array *Array, operationType ShiftType) *Array {
	result := array.Clone()
	result.ShiftRight(operationType)
	return result
}
