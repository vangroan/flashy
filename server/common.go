package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type errorPayload struct {
	ErrorMessage string `json:"errorMessage"`
}

// WriteJSONRequest is a helper function to write a JSON response
// to a given `http.ResponseWriter`.
func WriteJSONRequest(w http.ResponseWriter, statusCode int, payload interface{}) error {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// WriteJSONError is a helper function tp write a JSON error response
// to a given `http.ResponseWriter`.
func WriteJSONError(w http.ResponseWriter, logger *log.Entry, statusCode int, errPayload interface{}) error {
	var msg string
	if e, ok := errPayload.(error); ok {
		msg = e.Error()
	} else if s, ok := errPayload.(string); ok {
		msg = s
	} else {
		return fmt.Errorf("WriteJSONError(): Unknown type given in argument `err`")
	}

	logger.Warnf("WriteJSONError(): responding with error %d %s", statusCode, msg)

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(errorPayload{ErrorMessage: msg})
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}
