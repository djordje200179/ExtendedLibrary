package math

// SignedInteger is a constraint
// that requires the type to be a signed integer.
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInteger is a constraint
// that requires the type to be an unsigned integer.
type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint
// that requires the type to be an integer.
type Integer interface {
	SignedInteger | UnsignedInteger
}

// Float is a constraint
// that requires the type to be a floating point number.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint
// that requires the type to be a complex number.
type Complex interface {
	~complex64 | ~complex128
}

// Real is a constraint
// that requires the type to be a real number.
type Real interface {
	Integer | Float
}

// Number is a constraint
// that requires the type to be a number.
type Number interface {
	Real | Complex
}
