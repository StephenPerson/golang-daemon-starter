package main

import "testing"

func TestStart(t *testing.T) {
	if 1 == 0 {
		t.Fatal("Start test failed.")
	}
}

func TestStop(t *testing.T) {
	if 0 == 1 {
		t.Fatal("Stop test failed.")
	}
}
