package bitarray

import "math/bits"

// All returns true if all bits are set to 1.
func (array *Array) All() bool {
	sliceSize := len(array.slice)
	if array.lastElemOff != 0 {
		lastElementMask := uint8(0xFF) >> (8 - array.lastElemOff)
		lastElement := array.slice[sliceSize-1]

		if lastElement != lastElementMask {
			return false
		}

		sliceSize--
	}

	for i := range sliceSize - 1 {
		if array.slice[i] != 0xFF {
			return false
		}
	}

	return true
}

// Any returns true if any bit is set to 1.
func (array *Array) Any() bool {
	sliceSize := len(array.slice)
	if array.lastElemOff != 0 {
		lastElementMask := uint8(0xFF) >> (8 - array.lastElemOff)
		lastElement := array.slice[sliceSize-1]

		if lastElement&lastElementMask != 0 {
			return true
		}

		sliceSize--
	}

	for i := range sliceSize - 1 {
		if array.slice[i] != 0 {
			return true
		}
	}

	return false
}

// None returns true if no bit is set to 1.
func (array *Array) None() bool {
	sliceSize := len(array.slice)
	if array.lastElemOff != 0 {
		lastElementMask := uint8(0xFF) >> (8 - array.lastElemOff)
		lastElement := array.slice[sliceSize-1]

		if lastElement&lastElementMask != 0 {
			return false
		}

		sliceSize--
	}

	for i := range sliceSize - 1 {
		if array.slice[i] != 0 {
			return false
		}
	}

	return true
}

// Count returns the number of bits set to 1.
func (array *Array) Count() int {
	count := 0

	sliceSize := len(array.slice)
	if array.lastElemOff != 0 {
		lastElementMask := uint8(0xFF) >> (8 - array.lastElemOff)
		lastElement := array.slice[sliceSize-1]

		count += bits.OnesCount8(lastElement & lastElementMask)

		sliceSize--
	}

	for i := range sliceSize - 1 {
		elem := array.slice[i]

		count += bits.OnesCount8(elem)
	}

	return count
}
