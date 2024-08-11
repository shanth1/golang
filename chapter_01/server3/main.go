package main

import (
	"log"
	"net/http"

	lissajous "github.com/shanth1/the-go-programming-language/chapter_01/lissaious"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
