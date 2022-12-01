package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func modifiedNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "File not found", http.StatusNotFound)
}

func modifiedStripPrefix(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, errStat := os.Stat("." + r.URL.Path); errStat != nil {
			modifiedNotFound(w, r)
			return
		}
	})
}

func main() {
	fmt.Println("Launching server...")

	hh := http.HandlerFunc(helloHandler)
	http.Handle("/hello", hh)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", modifiedStripPrefix(fs))
	http.ListenAndServe(":12015", nil)
}
