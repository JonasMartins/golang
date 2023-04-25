package test

import "testing"

func TestOneEqualOne(t *testing.T) {
	got, want := 1, 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
