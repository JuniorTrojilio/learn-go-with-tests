package maps

import "errors"

// Dictionary is a group of words where we search for each word
type Dictionary map[string]string

// ErrNaoEncontrado is a Helper to error
var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra que você procura")

// Search is a method to abstract map tool
func (d Dictionary) Search(word string) (string, error) {
	def, exists := d[word]
	if !exists {
		return "", ErrNaoEncontrado
	}
	return def, nil
}

// Search is a func to abstract map tool
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func main() {

}
