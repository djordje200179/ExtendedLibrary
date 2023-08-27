package bitarray

func And(array1, array2 *Array) *Array {
	resArray := array1.Clone()
	resArray.And(array2)
	return resArray
}

func Or(array1, array2 *Array) *Array {
	resArray := array1.Clone()
	resArray.Or(array2)
	return resArray
}

func Xor(array1, array2 *Array) *Array {
	resArray := array1.Clone()
	resArray.Xor(array2)
	return resArray
}

func Not(array *Array) *Array {
	result := array.Clone()
	result.FlipAll()
	return result
}

func ShiftLeft(array *Array, operationType ShiftType) *Array {
	result := array.Clone()
	result.ShiftLeft(operationType)
	return result
}

func ShiftRight(array *Array, operationType ShiftType) *Array {
	result := array.Clone()
	result.ShiftRight(operationType)
	return result
}
