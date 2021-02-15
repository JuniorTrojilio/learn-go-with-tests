package main

import (
	"bytes"
	"testing"
)

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

func TestCount(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}

	Count(buffer, sleeperSpy)

	result := buffer.String()
	expect := `321Go!`

	if result != expect {
		t.Errorf("result '%s, expect '%s'", result, expect)
	}

	if sleeperSpy.Calls != 4 {
		t.Errorf("Don't have sufficient calls, expect 4, have %d", sleeperSpy.Calls)
	}
}
