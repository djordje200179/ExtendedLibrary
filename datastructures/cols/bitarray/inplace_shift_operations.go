package bitarray

// ShiftType is an enum for the type of shift operation to perform.
type ShiftType uint8

const (
	FillZero   ShiftType = iota // FillZero is a shift type that fills the empty bits with zeros.
	FillOne                     // FillOne is a shift type that fills the empty bits with ones.
	Arithmetic                  /*
		Arithmetic is a shift type that fills the empty bits with the same value
		as the sign bit during a right shift, and with zeros during a left shift.
	*/
)

// ShiftLeft performs an in-place left shift operation on the array.
func (array *Array) ShiftLeft(operationType ShiftType) {
	size := array.Size()
	if size == 0 {
		return
	}

	panic("not implemented")
}

// ShiftRight performs an in-place right shift operation on the array.
func (array *Array) ShiftRight(operationType ShiftType) {
	size := array.Size()
	if size == 0 {
		return
	}

	panic("not implemented")
}
