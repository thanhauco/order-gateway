package middleware
import (
    "net/http"
    "order-gateway/limiter"
)
var l = limiter.New()
func RateLimit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !l.Allow(r.RemoteAddr) {
            w.WriteHeader(429)
            return
        }
        next.ServeHTTP(w, r)
    })
}