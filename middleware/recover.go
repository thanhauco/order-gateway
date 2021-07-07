package middleware
import "log"
func Recover(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Println("Panic:", err)
                w.WriteHeader(500)
            }
        }()
        next.ServeHTTP(w, r)
    })
}