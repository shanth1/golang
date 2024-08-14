package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		cyclesStr := queryParams.Get("cycles")
		var cycles int
		var err error

		if cyclesStr != "" {
			cycles, err = strconv.Atoi(cyclesStr)
			if err != nil {
				http.Error(w, "Invalid cycles parameter", http.StatusBadRequest)
				return
			}
		}

		opts := &Options{
			Cycles: cycles,
		}
		lissajous(w, opts)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
