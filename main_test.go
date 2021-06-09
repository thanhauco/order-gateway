package main
import "go.uber.org/goleak"
import "testing"
func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}