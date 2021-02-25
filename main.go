package main
import (
    "net/http"
    "order-gateway/config"
    "order-gateway/proxy"
    "order-gateway/lb"
    "order-gateway/middleware"
    "order-gateway/handlers"
    "fmt"
    "os"
    "os/signal"
    "context"
    "time"
)
func main() {
    c := config.Load()
    l := lb.New([]string{c.Target})
    p := proxy.New(l)
    h := middleware.RateLimit(middleware.Measure(middleware.Logger(middleware.Auth(p))))
    mux := http.NewServeMux()
    mux.Handle("/", h)
    mux.HandleFunc("/metrics", handlers.Stats)
    srv := &http.Server{Addr: fmt.Sprintf(":%d", c.Port), Handler: mux}
    go func() {
        if err := srv.ListenAndServe(); err != nil { fmt.Println(err) }
    }()
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    <-stop
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    srv.Shutdown(ctx)
}