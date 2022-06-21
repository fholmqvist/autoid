package autoid

import (
	"math"
	"testing"
)

func Test8(t *testing.T) {
	defer Reset()
	if id, err := NextUint8(); id != 1 || err != nil {
		t.Fatalf("expected (1, <nil>), got (%d, %v)", id, err)
	}
	Set(math.MaxUint8)
	if _, err := NextUint8(); err == nil {
		t.Fatalf("expected overflow, __id was %d", __id)
	}
}

func Test16(t *testing.T) {
	defer Reset()
	if id, err := NextUint16(); id != 1 || err != nil {
		t.Fatalf("expected (1, <nil>), got (%d, %v)", id, err)
	}
	Set(math.MaxUint16)
	if _, err := NextUint16(); err == nil {
		t.Fatalf("expected overflow, __id was %d", __id)
	}
}
func Test32(t *testing.T) {
	defer Reset()
	if id, err := NextUint32(); id != 1 || err != nil {
		t.Fatalf("expected (1, <nil>), got (%d, %v)", id, err)
	}
	Set(math.MaxUint32)
	if _, err := NextUint32(); err == nil {
		t.Fatalf("expected overflow, __id was %d", __id)
	}
}

func Test64(t *testing.T) {
	defer Reset()
	if id, err := NextUint64(); id != 1 || err != nil {
		t.Fatalf("expected (1, <nil>), got (%d, %v)", id, err)
	}
	Set(math.MaxUint64)
	if _, err := NextUint64(); err == nil {
		t.Fatalf("expected overflow, __id was %d", __id)
	}
}
