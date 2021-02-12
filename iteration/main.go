package iteration

// Repeat is a loop for print caracter
func Repeat(caracter string) string {
	var repetitions string

	for i := 0; i < 5; i++ {
		repetitions += caracter
	}

	return repetitions
}

func main() {

}
