package middleware
import (
    "net/http"
    "order-gateway/metrics"
    "sync/atomic"
)
func Measure(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        atomic.AddUint64(&metrics.Global.Requests, 1)
        next.ServeHTTP(w, r)
    })
}