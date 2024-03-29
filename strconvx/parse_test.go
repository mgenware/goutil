package strconvx

import (
	"fmt"
	"math"
	"testing"

	"github.com/mgenware/goutil/test"
)

func TestParseFloat64(t *testing.T) {
	r, err := ParseFloat64("32")
	test.FatalOn(err, t)
	test.Assert(t, r, float64(32))

	r, err = ParseFloat64("-32.003")
	test.FatalOn(err, t)
	test.Assert(t, r, float64(-32.003))

	r, err = ParseFloat64("aaa")
	if err == nil {
		t.Fatal("Expected error")
	}
	test.Assert(t, r, float64(0))

	r, err = ParseFloat64(fmt.Sprintf("%v", math.MaxFloat64))
	test.FatalOn(err, t)
	test.Assert(t, r, math.MaxFloat64)
}

func TestParseFloat32(t *testing.T) {
	if r, _ := ParseFloat32("32"); r != float32(32) {
		t.Errorf("Expected %v, got %v", float32(32), r)
	}
	if r, _ := ParseFloat32("-32.003"); r != -32.003 {
		t.Errorf("Expected %v, got %v", float32(-32.003), r)
	}
	if _, err := ParseFloat32("aaa"); err == nil {
		t.Error("Error expected")
	}
	if _, err := ParseFloat32(fmt.Sprintf("%v", math.MaxFloat64)); err == nil {
		t.Error("Error expected")
	}
	if r, _ := ParseFloat32(fmt.Sprintf("%v", math.MaxFloat32)); r != math.MaxFloat32 {
		t.Errorf("Expected %v, got %v", math.MaxFloat32, r)
	}
}

func TestParseInt(t *testing.T) {
	if r, _ := ParseInt("32"); r != 32 {
		t.Errorf("Expected %v, got %v", int(32), r)
	}
	if r, _ := ParseInt("-3"); r != -3 {
		t.Errorf("Expected %v, got %v", int(-3), r)
	}
	if _, err := ParseInt("aaa"); err == nil {
		t.Error("Error expected")
	}
}

func TestParseUint(t *testing.T) {
	if r, _ := ParseUint("32"); r != 32 {
		t.Errorf("Expected %v, got %v", 32, r)
	}
	if _, err := ParseUint("-32"); err == nil {
		t.Error("Error expected")
	}
	if _, err := ParseUint("aaa"); err == nil {
		t.Error("Error expected")
	}
}

func TestParseInt64(t *testing.T) {
	if r, _ := ParseInt64("-9223372036854775808"); r != int64(-9223372036854775808) {
		t.Errorf("Expected %v, got %v", int64(-9223372036854775808), r)
	}
	if r, _ := ParseInt64("-9223372036854775808"); r != int64(-9223372036854775808) {
		t.Errorf("Expected %v, got %v", int64(-9223372036854775808), r)
	}
	if _, err := ParseInt64("18446744073709551615"); err == nil {
		t.Error("Error expected")
	}
}

func TestParseUint64(t *testing.T) {
	if r, _ := ParseUint64("18446744073709551615"); r != uint64(18446744073709551615) {
		t.Errorf("Expected %v, got %v", uint64(18446744073709551615), r)
	}
	if _, err := ParseUint64("-32"); err == nil {
		t.Error("Error expected")
	}
}
