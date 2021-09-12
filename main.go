// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	titleCollection := [...]string{"First Title", "Second Title", "Third Title"}

	for position, title := range titleCollection {
		fmt.Fprintf(w, "%d: %s ", position, title)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
