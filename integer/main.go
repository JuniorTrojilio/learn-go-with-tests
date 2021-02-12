package integer

import "fmt"

// Sum is a sum of two numbers
func Sum(numberOne, numberTwo int) int {
	return numberOne + numberTwo
}

// ExampleSum is a example of Sum function
func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func main() {

}
