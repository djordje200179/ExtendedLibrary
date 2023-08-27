package bitarray

type ShiftType uint8

const (
	FillZero ShiftType = iota
	FillOne
	Logical
	Arithmetic
)

func (array *Array) ShiftLeft(operationType ShiftType) {
	size := array.Size()
	if size == 0 {
		return
	}

	panic("not implemented")
}

func (array *Array) ShiftRight(operationType ShiftType) {
	size := array.Size()
	if size == 0 {
		return
	}

	panic("not implemented")
}
