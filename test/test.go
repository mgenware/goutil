package test

import (
	"runtime/debug"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Assert compares the given two values, and make testing fail if they are not equivalent.
func Assert(t *testing.T, got, want interface{}) {
	if diff := cmp.Diff(want, got); diff != "" {
		debug.PrintStack()
		t.Fatalf("Mismatch (-want +got):\n%s", diff)
	}
}

// PanicOn panics if the given error is not nil.
func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

// FatalOn calls t.Fatal if the given error is not nil.
func FatalOn(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
