package main
import (
    "net/http"
    "order-gateway/config"
    "order-gateway/proxy"
    "order-gateway/lb"
    "order-gateway/middleware"
    "fmt"
)
func main() {
    c := config.Load()
    l := lb.New([]string{c.Target})
    p := proxy.New(l)
    h := middleware.Logger(middleware.Auth(p))
    http.Handle("/", h)
    fmt.Printf("Listening on %d\n", c.Port)
    http.ListenAndServe(fmt.Sprintf(":%d", c.Port), nil)
}