package main
import "order-gateway/handlers"
// ... setup
http.HandleFunc("/health", handlers.Health)