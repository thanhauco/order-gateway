package middleware
import "github.com/google/uuid"
func RequestID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Request-ID", uuid.New().String())
        next.ServeHTTP(w, r)
    })
}