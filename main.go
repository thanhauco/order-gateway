package main
import (
    "net/http"
    "order-gateway/config"
    "order-gateway/proxy"
    "fmt"
)
func main() {
    c := config.Load()
    p := proxy.New(c.Target)
    http.Handle("/", p)
    fmt.Printf("Listening on %d\n", c.Port)
    http.ListenAndServe(fmt.Sprintf(":%d", c.Port), nil)
}