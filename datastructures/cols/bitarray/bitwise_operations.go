package bitarray

// And performs bitwise AND operation on two arrays and returns the result.
func And(array1, array2 *Array) *Array {
	resArray := array1.Clone()
	resArray.And(array2)
	return resArray
}

// Or performs bitwise OR operation on two arrays and returns the result.
func Or(array1, array2 *Array) *Array {
	resArray := array1.Clone()
	resArray.Or(array2)
	return resArray
}

// Xor performs bitwise XOR operation on two arrays and returns the result.
func Xor(array1, array2 *Array) *Array {
	resArray := array1.Clone()
	resArray.Xor(array2)
	return resArray
}

// Not performs bitwise NOT operation on an array and returns the result.
func Not(array *Array) *Array {
	result := array.Clone()
	result.FlipAll()
	return result
}

// ShiftLeft performs bitwise left shift operation on an array and returns the result.
func ShiftLeft(array *Array, operationType ShiftType) *Array {
	result := array.Clone()
	result.ShiftLeft(operationType)
	return result
}

// ShiftRight performs bitwise right shift operation on an array and returns the result.
func ShiftRight(array *Array, operationType ShiftType) *Array {
	result := array.Clone()
	result.ShiftRight(operationType)
	return result
}
