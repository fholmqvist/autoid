package autoid

import (
	"fmt"
	"math"
)

var (
	__id        uint64 = 0
	errOverflow string = "overflow for %v: %d"
)

func Next() (uint32, error) {
	return NextUint32()
}

func Set(id uint64) {
	__id = id
}

func Reset() {
	__id = 0
}

func NextUint8() (uint8, error) {
	id := next()
	if id > math.MaxUint8 {
		return 0, fmt.Errorf(errOverflow, "uint8", id)
	}
	return uint8(id), nil
}

func NextUint16() (uint16, error) {
	id := next()
	if id > math.MaxUint16 {
		return 0, fmt.Errorf(errOverflow, "uint16", id)
	}
	return uint16(id), nil
}

func NextUint32() (uint32, error) {
	id := next()
	if id > math.MaxUint32 {
		return 0, fmt.Errorf(errOverflow, "uint32", id)
	}
	return uint32(id), nil
}

func NextUint64() (uint64, error) {
	if __id == math.MaxUint64 {
		return 0, fmt.Errorf(errOverflow, "uint32", __id)
	}
	return next(), nil
}

func next() uint64 {
	__id++
	return __id
}
