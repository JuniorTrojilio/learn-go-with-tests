package main

import "fmt"

// Hello prints "hello" + name
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Junior"))
}
