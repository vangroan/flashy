package main

import (
	"context"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.TraceLevel)
	log.Debug("main(): Start")

	var wg sync.WaitGroup

	r := mux.NewRouter()
	r.Use(logRequest)
	NewFlashCardCtrl(r.PathPrefix("/cards").Subrouter(), nil)

	srv := http.Server{Addr: ":8000"}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// returns ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		} else {
			log.Debug("ListenAndServe(): Graceful shutdown")
		}
	}()

	log.Debug("main(): Listening")

	// In real world you don't use TODO()
	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	log.Debug("main(): Shutting down")
	wg.Wait()
}
