package lb
import (
    "net/url"
    "sync/atomic"
)
type LoadBalancer struct {
    targets []*url.URL
    current uint64
}
func New(targets []string) *LoadBalancer {
    var urls []*url.URL
    for _, t := range targets {
        u, _ := url.Parse(t)
        urls = append(urls, u)
    }
    return &LoadBalancer{targets: urls}
}
func (l *LoadBalancer) Next() *url.URL {
    i := atomic.AddUint64(&l.current, 1)
    return l.targets[i % uint64(len(l.targets))]
}