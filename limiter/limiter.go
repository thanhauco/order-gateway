package limiter
import "sync"
type Limiter struct {
    tokens map[string]int
    mu sync.Mutex
}
func New() *Limiter { return &Limiter{tokens: make(map[string]int)} }
func (l *Limiter) Allow(ip string) bool {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.tokens[ip]++
    return l.tokens[ip] < 100
}