package middleware
import "compress/gzip"
import "io"
import "net/http"
type gzipWriter struct {
    http.ResponseWriter
    Writer io.Writer
}
func (w gzipWriter) Write(b []byte) (int, error) { return w.Writer.Write(b) }
func Gzip(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        gw := gzip.NewWriter(w)
        defer gw.Close()
        next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gw}, r)
    })
}