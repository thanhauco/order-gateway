package middleware
import (
    "net/http"
    "go.opentelemetry.io/otel"
)
func Trace(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx, span := otel.Tracer("gateway").Start(r.Context(), r.URL.Path)
        defer span.End()
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}