package main

import (
	"fmt"
	"net/http"
)

// NotFound is the default handler for unknown routes.
func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "404 not found")
}
