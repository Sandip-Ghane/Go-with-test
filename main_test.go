package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("test to hello people", func(t *testing.T) {
		got := Hello("sandip")
		want := "Hello sandip"
		if got != want {
			t.Errorf("Expect: %s but got: %s", want, got)
		}
	})
	t.Run("test to empty people name", func(t *testing.T) {
		got := Hello("")
		want := "Hello world"
		if got != want {
			t.Errorf("Expect: %s but got: %s", want, got)
		}
	})
}
