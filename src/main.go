package main
import (
    "context"
    "os/signal"
    "syscall"
    "os"
)
func main() {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()
    <-ctx.Done()
}