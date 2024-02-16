package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DaCondor54/roadmapsh-backend/internal/handlers"
	"github.com/joho/godotenv"
)

const keyServerAddr = "serverAddr"

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln("Erro Loading .env file")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /auth", handlers.Authenticate)
	mux.HandleFunc("POST /", handlers.PostRoot)
	mux.HandleFunc("GET /", handlers.GetRoot)

	server1 := &http.Server{
		Addr:    ":3333",
		Handler: mux,
	}

	fmt.Println("Starting Server 1")
	err = server1.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server 1 Closed")
	} else if err != nil {
		fmt.Printf("Server 1 Error %s", err)
		os.Exit(1)
	}
}
