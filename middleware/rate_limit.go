package middleware

import (
	"net/http"
	"time"
)

var requestCounts = make(map[string]int)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		requestCounts[clientIP]++

		if requestCounts[clientIP] > 100 { // Limit to 100 requests
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		time.AfterFunc(time.Minute, func() { requestCounts[clientIP]-- })
		next.ServeHTTP(w, r)
	})
}
