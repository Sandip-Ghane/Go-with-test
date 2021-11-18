package array

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		numbers := [...]int{1, 2, 3, 4, 5}
		got := SumN(numbers)
		want := 15
		if got != want {
			t.Errorf("Expected: %v but got: %v", got, want)
		}
	})
}
