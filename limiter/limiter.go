package limiter
import "sync"
type Limiter struct {
    tokens map[string]int
    mu sync.Mutex
}
func New() *Limiter { return &Limiter{tokens: make(map[string]int)} }