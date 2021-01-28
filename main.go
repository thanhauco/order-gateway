package main
import (
    "net/http"
    "order-gateway/config"
    "order-gateway/proxy"
    "order-gateway/lb"
    "fmt"
)
func main() {
    c := config.Load()
    l := lb.New([]string{c.Target, "http://backup:9000"})
    p := proxy.New(l)
    http.Handle("/", p)
    fmt.Printf("Listening on %d\n", c.Port)
    http.ListenAndServe(fmt.Sprintf(":%d", c.Port), nil)
}