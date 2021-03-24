package main
import (
    "net/http"
    "order-gateway/config"
    // ... imports
)
func main() {
    // ... setup
    server := &http.Server{
        Addr: ":8443",
        TLSConfig: config.LoadTLS(),
    }
    server.ListenAndServeTLS("cert.pem", "key.pem")
}