package middleware
import "go.uber.org/zap"
var logger, _ = zap.NewProduction()
func ZapLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        logger.Info("request", zap.String("path", r.URL.Path))
        next.ServeHTTP(w, r)
    })
}