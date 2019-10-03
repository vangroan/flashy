package main

import (
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "404 not found")
}
