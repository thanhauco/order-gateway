package proxy
import (
    "net/http"
    "net/http/httputil"
    "order-gateway/lb"
)
func New(l *lb.LoadBalancer) http.Handler {
    return &httputil.ReverseProxy{
        Director: func(req *http.Request) {
            target := l.Next()
            req.URL.Scheme = target.Scheme
            req.URL.Host = target.Host
        },
    }
}