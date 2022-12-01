package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Launching server...")

	http.ListenAndServe(":12015", http.FileServer(http.Dir(".")))
}
