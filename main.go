package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/frequency", handleStringProcessing)
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		return
	}
}
