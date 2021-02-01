package metrics
import "sync/atomic"
type Metrics struct {
    Requests uint64
    Errors uint64
}
var Global Metrics