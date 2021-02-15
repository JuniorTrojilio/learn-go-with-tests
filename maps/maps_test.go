package maps

import (
	"testing"
)

func compareStrings(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("Result '%s', expect '%s', dado '%s'", result, expect, "test")
	}

}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is only a test"}

	t.Run("Know word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expect := "This is only a test"

		compareStrings(t, result, expect)
	})

	t.Run("Unknown words", func(t *testing.T) {
		_, err := dictionary.Search("unknow")

		if err == nil {
			t.Fatal("Expect that obtain a error.")
		}
	})
}
