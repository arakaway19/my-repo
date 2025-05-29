package main

import "testing"

func TestEven0r0ad(t *testing.T) {
	result := Even0r0ad(10)
	if result != "even" {
		t.Errorf("expected 'even', actual: %s", result)
	}
}