package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup Logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.TraceLevel)
	log.Debug("main(): Start")

	var wg sync.WaitGroup

	err := os.MkdirAll("./data", os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Database
	db, err := gorm.Open("sqlite3", "./data/data.sqlite3")
	if err != nil {
		panic(err)
	}

	// Migrate
	db.AutoMigrate(&Box{}, &BoxCard{}, &FlashCard{})

	// Routes
	r := mux.NewRouter()
	r.Use(logRequest)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	NewFlashCardCtrl(r.PathPrefix("/cards").Subrouter(), db)

	srv := http.Server{Addr: ":8000"}
	srv.Handler = r

	// Wait for listen in goroutine
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

	// Custom OS signal handling
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for s := range sig {
			log.Debugf("main(): Received signal %s", s.String())

			// In real world you don't use TODO()
			if err := srv.Shutdown(context.TODO()); err != nil {
				panic(err) // failure/timeout shutting down the server gracefully
			}

			return
		}
	}()

	log.Debug("main(): Shutting down")
	wg.Wait()
}
