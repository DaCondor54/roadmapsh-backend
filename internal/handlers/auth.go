package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": "some-more-info",
	})

	value := os.Getenv("SECRET")
	fmt.Println(value)
	tokenString, err := token.SignedString([]byte(value))
	if err != nil {
		fmt.Fprintf(w, "Failed to generate jwt\n")
	}
	fmt.Fprintf(w, tokenString)
}
