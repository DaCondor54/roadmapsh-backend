package handlers

import (
	"fmt"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value := ctx.Value("serverAddr")
	fmt.Printf("get / was called from %s\n", value)
	fmt.Fprintf(w, "Hello from %s\n", value)
}
