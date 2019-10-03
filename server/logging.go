package main

import (
	"context"
	b64 "encoding/base64"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const (
	operationIDLogKey string = "op"
)

// logRequest is middleware that will add a correlation
// ID and primed logger to the request context.
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		opID := generateOperationID()
		logger := log.WithFields(log.Fields{
			operationIDLogKey: opID,
		})
		logger.Debug("logRequest(): Request")

		ctx1 := context.WithValue(r.Context(), operationIDKey, opID)
		ctx2 := context.WithValue(ctx1, loggerKey, logger)

		req := r.WithContext(ctx2)

		next.ServeHTTP(w, req)
	})
}

func generateOperationID() string {
	id, _ := uuid.NewUUID()
	bin, _ := id.MarshalBinary()
	return b64.URLEncoding.EncodeToString(bin)
}
