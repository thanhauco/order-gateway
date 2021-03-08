package middleware
import (
    "net/http"
    "github.com/afex/hystrix-go/hystrix"
)
func CircuitBreaker(next http.Handler) http.Handler {
    hystrix.ConfigureCommand("api", hystrix.CommandConfig{Timeout: 1000})
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        err := hystrix.Do("api", func() error {
            next.ServeHTTP(w, r)
            return nil
        }, nil)
        if err != nil { w.WriteHeader(503) }
    })
}