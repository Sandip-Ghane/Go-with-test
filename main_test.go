package main

import "testing"

func TestMain(t *testing.T) {
	got := Hello()
	want := "Hello World"
	if got != "Hello World" {
		t.Errorf("Expect: %s but got: %s", want, got)
	}
}
