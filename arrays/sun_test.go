package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	/*t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if want != got {
			t.Errorf("Want %d but got %d, given %v", want, got, numbers)
		}
	})
	*/
	t.Run("Collection of any nsize", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if want != got {
			t.Errorf("Want %d but got %d, given %v", want, got, numbers)
		}
	})
}
