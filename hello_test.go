package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Junior")
	expect := "Hello, Junior"

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}

}
