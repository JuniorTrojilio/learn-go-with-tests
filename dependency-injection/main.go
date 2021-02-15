package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet peoples
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func handlerMyGreet(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(handlerMyGreet))

	if err != nil {
		fmt.Println(err)
	}
}
