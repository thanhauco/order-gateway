package main
import "time"
// ... existing
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)