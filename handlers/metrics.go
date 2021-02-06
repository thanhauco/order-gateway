package handlers
import (
    "net/http"
    "order-gateway/metrics"
    "fmt"
    "sync/atomic"
)
func Stats(w http.ResponseWriter, r *http.Request) {
    reqs := atomic.LoadUint64(&metrics.Global.Requests)
    fmt.Fprintf(w, "Requests: %d", reqs)
}