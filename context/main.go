package main

import (
	"context"
	"fmt"
	"net/http"
)

// Store is a interface to store object
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server is a server
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())
		fmt.Fprint(w, data)
	}
}

func main() {

}
