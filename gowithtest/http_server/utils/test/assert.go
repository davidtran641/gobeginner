package test

import (
	"reflect"
	"testing"
)

// AssertEqual ...
func AssertEqual(t *testing.T, want interface{}, got interface{}) {
	t.Helper()
	if reflect.DeepEqual(want, got) {
		return
	}

	t.Errorf("want %v, but got %v", want, got)
}
