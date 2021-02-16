package lb
import "testing"
func TestNext(t *testing.T) {
    l := New([]string{"a", "b"})
    if l.Next().String() != "a" { t.Fail() }
    if l.Next().String() != "b" { t.Fail() }
}