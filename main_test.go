package main

import "testing"

func TestMain(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Expected: %q but want: %q", got, want)
		}
	}
	t.Run("test to hello people", func(t *testing.T) {
		got := Hello("Sandip")
		want := "Hello Sandip"
		assertCorrectMessage(t, got, want)
	})
	t.Run("test to empty people", func(t *testing.T) {
		got := Hello("")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
}
