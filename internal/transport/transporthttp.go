package transport

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type requestIDKey struct{}

func NewMux(h Handler) *mux.Router {
	m := mux.NewRouter()

	m.Use(requestIDMiddleWare)

	m.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	return m
}

func requestIDMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()

		// Set the request ID in the request context
		ctx := context.WithValue(r.Context(), requestIDKey{}, requestID)

		// Optionally, set the request ID in the response header
		w.Header().Set("X-Request-ID", requestID)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
