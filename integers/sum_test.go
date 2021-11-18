package integers

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got := Sum(2, 2)
		want := 4
		if got != want {
			t.Errorf("Expected: %d but got: %d", want, got)
		}
	})
}
