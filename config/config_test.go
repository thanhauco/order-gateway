package config
import "testing"
func TestLoad(t *testing.T) {
    c := Load()
    if c.Port != 8080 { t.Fail() }
}