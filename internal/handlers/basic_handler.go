package handlers

import (
	"fmt"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get / was called\n")
	fmt.Fprintf(w, "Hello from GET\n")
}

func PostRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("psot / was called\n")
	fmt.Fprintf(w, "Hello from POST\n")
}
