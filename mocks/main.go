package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper is a interface to Sleep
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is a default sleeper
type DefaultSleeper struct{}

// Sleep is a method to abstract truly Sleep
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Count up to 3
func Count(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprint(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, "Go!")
}

func main() {
	sleeper := &DefaultSleeper{}
	Count(os.Stdout, sleeper)
}
