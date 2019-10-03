package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RestCreate is a controller that creates a resource
type RestCreate interface {
	Create(w http.ResponseWriter, r *http.Request)
}

// RestRead is a controller that reads a single resource
type RestRead interface {
	Read(w http.ResponseWriter, r *http.Request)
}

// RestReadMany is a controller that reads multiple resources
type RestReadMany interface {
	ReadMany(w http.ResponseWriter, r *http.Request)
}

// RestUpdate is a controller that replaces a resource
type RestUpdate interface {
	Update(w http.ResponseWriter, r *http.Request)
}

// RestPartialUpdate is a controller that mutates a resource
type RestPartialUpdate interface {
	PartialUpdate(w http.ResponseWriter, r *http.Request)
}

// RestDelete is a controller that removes a resource
type RestDelete interface {
	Delete(w http.ResponseWriter, r *http.Request)
}

// RegisterController adds all the resources endpoints found
// on the given controller to the provided route.
func RegisterController(ctrl interface{}, route *mux.Router) {
	if handler, ok := ctrl.(RestCreate); ok {
		route.Methods("POST").Path("/").Handler(http.HandlerFunc(handler.Create))
	}

	if handler, ok := ctrl.(RestRead); ok {
		route.Methods("GET").Path("/{id}").Handler(http.HandlerFunc(handler.Read))
	}

	if handler, ok := ctrl.(RestReadMany); ok {
		route.Methods("GET").Path("/").Handler(http.HandlerFunc(handler.ReadMany))
	}

	if handler, ok := ctrl.(RestUpdate); ok {
		route.Methods("PUT").Path("/{id}").Handler(http.HandlerFunc(handler.Update))
	}

	if handler, ok := ctrl.(RestPartialUpdate); ok {
		route.Methods("PATCH").Path("/{id}").Handler(http.HandlerFunc(handler.PartialUpdate))
	}

	if handler, ok := ctrl.(RestDelete); ok {
		route.Methods("DELETE").Path("/{id}").Handler(http.HandlerFunc(handler.Delete))
	}
}
