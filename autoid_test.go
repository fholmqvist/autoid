package autoid

import (
	"math"
	"testing"
)

func Test8(t *testing.T) {
	defer Reset()
	if n, err := NextUint8(); n != 1 || err != nil {
		t.Fatalf("expected (1, <nil>), got (%d, %v)", n, err)
	}
	id = math.MaxUint8
	if _, err := NextUint8(); err == nil {
		t.Fatalf("expected overflow, got %d", id)
	}
}

func Test16(t *testing.T) {
	defer Reset()
	if n, err := NextUint16(); n != 1 || err != nil {
		t.Fatalf("expected (1, <nil>), got (%d, %v)", n, err)
	}
	id = math.MaxUint16
	if _, err := NextUint16(); err == nil {
		t.Fatalf("expected overflow, got %d", id)
	}
}
func Test32(t *testing.T) {
	defer Reset()
	if n, err := NextUint32(); n != 1 || err != nil {
		t.Fatalf("expected (1, <nil>), got (%d, %v)", n, err)
	}
	id = math.MaxUint32
	if _, err := NextUint32(); err == nil {
		t.Fatalf("expected overflow, got %d", id)
	}
}

func Test64(t *testing.T) {
	defer Reset()
	if n, err := NextUint64(); n != 1 || err != nil {
		t.Fatalf("expected (1, <nil>), got (%d, %v)", n, err)
	}
	id = math.MaxUint64
	if _, err := NextUint64(); err == nil {
		t.Fatalf("expected overflow, got %d", id)
	}
}
