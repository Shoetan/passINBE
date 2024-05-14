package middleware

import (
	"net/http"
)

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Allow requests from any origin
			w.Header().Set("Access-Control-Allow-Origin", "*")

			// Set the allowed methods for CORS requests
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")

			// Set the allowed headers for CORS requests
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			// Allow cookies to be sent and received in CORS requests
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			// If it's a preflight request, respond with 200 OK
			if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusOK)
					return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
	})
}