package main

import "testing"

func TestAdder(t *testing.T) {
	if 1 == 2 {
		t.Fail()
	}
}
