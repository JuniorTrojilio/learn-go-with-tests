package hello

import "testing"

func TestHello(t *testing.T) {

	verifyCorrectMessage := func(t *testing.T, result, expect string) {
		t.Helper()
		if result != expect {
			t.Errorf("result '%s', expect '%s'", result, expect)
		}
	}

	t.Run("Sea hello to specified people", func(t *testing.T) {
		name := "Junior"
		result := Hello(name)
		expect := "Hello, " + name

		verifyCorrectMessage(t, result, expect)
	})

	t.Run("Sea hello to all world", func(t *testing.T) {
		result := Hello("")
		expect := "Hello, world"

		verifyCorrectMessage(t, result, expect)
	})

}
