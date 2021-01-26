package middleware
import "net/http"
func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("Authorization") == "" {
            w.WriteHeader(401)
            return
        }
        next.ServeHTTP(w, r)
    })
}