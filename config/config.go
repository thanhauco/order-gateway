package config
import (
    "os"
    "strconv"
)
type Config struct {
    Port int
    Target string
}
func Load() *Config {
    p, _ := strconv.Atoi(os.Getenv("PORT"))
    if p == 0 { p = 8080 }
    t := os.Getenv("TARGET")
    if t == "" { t = "http://localhost:9000" }
    return &Config{Port: p, Target: t}
}