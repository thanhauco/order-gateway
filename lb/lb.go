package lb
import "net/url"
type LoadBalancer struct {
    targets []*url.URL
    current uint64
}