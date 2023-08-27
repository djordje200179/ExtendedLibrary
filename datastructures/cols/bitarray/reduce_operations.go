package bitarray

import "math/bits"

func (array *Array) All() bool {
	sliceSize := array.array.Size()
	if array.lastElemOff != 0 {
		lastElementMask := uint8(0xFF) >> (8 - array.lastElemOff)
		lastElement := array.array.Get(sliceSize - 1)

		if lastElement != lastElementMask {
			return false
		}

		sliceSize--
	}

	for i := 0; i < sliceSize-1; i++ {
		if array.array.Get(i) != 0xFF {
			return false
		}
	}

	return true
}

func (array *Array) Any() bool {
	sliceSize := array.array.Size()
	if array.lastElemOff != 0 {
		lastElementMask := uint8(0xFF) >> (8 - array.lastElemOff)
		lastElement := array.array.Get(sliceSize - 1)

		if lastElement&lastElementMask != 0 {
			return true
		}

		sliceSize--
	}

	for i := 0; i < sliceSize-1; i++ {
		if array.array.Get(i) != 0 {
			return true
		}
	}

	return false
}

func (array *Array) None() bool {
	sliceSize := array.array.Size()
	if array.lastElemOff != 0 {
		lastElementMask := uint8(0xFF) >> (8 - array.lastElemOff)
		lastElement := array.array.Get(sliceSize - 1)

		if lastElement&lastElementMask != 0 {
			return false
		}

		sliceSize--
	}

	for i := 0; i < sliceSize-1; i++ {
		if array.array.Get(i) != 0 {
			return false
		}
	}

	return true
}

func (array *Array) Count() int {
	count := 0

	sliceSize := array.array.Size()
	if array.lastElemOff != 0 {
		lastElementMask := uint8(0xFF) >> (8 - array.lastElemOff)
		lastElement := array.array.Get(sliceSize - 1)

		count += bits.OnesCount8(lastElement & lastElementMask)

		sliceSize--
	}

	for i := 0; i < sliceSize-1; i++ {
		elem := array.array.Get(i)

		count += bits.OnesCount8(elem)
	}

	return count
}
