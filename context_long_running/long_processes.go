package context_long_running

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// send context , return response
		// not care about error
		data, _ := store.Fetch(r.Context())
		fmt.Fprint(w, data)
	}
}
