package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Junior")
	result := buffer.String()
	expect := "Hello, Junior"

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}
}
