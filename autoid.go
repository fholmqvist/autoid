package autoid

import (
	"fmt"
	"math"
)

var (
	id          uint64 = 0
	errOverflow string = "overflow for %v: %d"
)

func Reset() {
	id = 0
}

func Next() (uint32, error) {
	return NextUint32()
}

func NextUint8() (uint8, error) {
	n := next()
	if n > math.MaxUint8 {
		return 0, fmt.Errorf(errOverflow, "uint8", n)
	}
	return uint8(n), nil
}

func NextUint16() (uint16, error) {
	n := next()
	if n > math.MaxUint16 {
		return 0, fmt.Errorf(errOverflow, "uint16", n)
	}
	return uint16(n), nil
}

func NextUint32() (uint32, error) {
	n := next()
	if n > math.MaxUint32 {
		return 0, fmt.Errorf(errOverflow, "uint32", n)
	}
	return uint32(n), nil
}

func NextUint64() (uint64, error) {
	if id == math.MaxUint64 {
		return 0, fmt.Errorf(errOverflow, "uint32", id)
	}
	return next(), nil
}

func next() uint64 {
	id++
	return id
}
