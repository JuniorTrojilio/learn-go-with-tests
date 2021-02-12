package integer

import "testing"

func TestSum(t *testing.T) {

	verifyCorrectMessage := func(t *testing.T, result, expect int) {
		t.Helper()
		if result != expect {
			t.Errorf("result '%d', expect '%d'", result, expect)
		}
	}

	result := Sum(2, 2)
	expect := 4

	verifyCorrectMessage(t, result, expect)
}
