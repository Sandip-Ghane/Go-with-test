package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeated", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		if got != want {
			t.Errorf("Expected: %q but want: %q", got, want)
		}
	})
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("a")
	}
}
