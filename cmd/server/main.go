package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/DaCondor54/roadmapsh-backend/internal/handlers"
)

const keyServerAddr = "serverAddr"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.GetRoot)
	mux.HandleFunc("POST /", handlers.PostRoot)
	ctx, cancelCtx := context.WithCancel(context.Background())

	server1 := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	server2 := http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	go func() {
		fmt.Println("Starting Server 1")
		err := server1.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server 1 Closed")
		} else if err != nil {
			fmt.Printf("Server 1 Error %s", err)
		}
		cancelCtx()
	}()

	go func() {
		fmt.Println("Starting Server 2")
		err := server2.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server 2 Closed")
		} else if err != nil {
			fmt.Printf("Server 2 Error %s", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()

}
