package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value := ctx.Value("serverAddr")
	fmt.Printf("get / was called from %s\n", value)
	fmt.Fprintf(w, "Hello from %s\n", value)
}

func PostRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value := ctx.Value("serverAddr")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid Body %s", err)
		return
	}

	fmt.Printf("post / was called from %s with %s\n", value, body)
	fmt.Fprintf(w, "post / was called from %s with %s\n", value, body)

}
