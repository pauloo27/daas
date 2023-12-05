package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("text")
		rawMaxSize := r.URL.Query().Get("maxSize")

		maxSize, err := strconv.Atoi(rawMaxSize)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("maxSize must be a number\n"))
			return
		}

		if len(text) > maxSize {
			w.Write([]byte(text[:maxSize] + "…"))
		} else {
			w.Write([]byte(text))
		}

		w.Write([]byte{'\n'})
	})

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}
