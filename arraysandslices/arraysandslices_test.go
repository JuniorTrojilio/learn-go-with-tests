package arraysandslices

import "testing"

func TestSoma(t *testing.T) {
	t.Run("Sum colection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expect := 15

		if expect != result {
			t.Errorf("expect '%d', but received '%d' dado '%v'", expect, result, numbers)
		}
	})
}
