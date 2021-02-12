package arraysandslices

// Sum returns a sum of int arrays
func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func main() {

}
