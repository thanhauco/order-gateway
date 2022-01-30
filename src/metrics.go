package main
import "github.com/prometheus/client_golang/prometheus"
var requests = prometheus.NewCounter(prometheus.CounterOpts{
    Name: "requests_total",
    Help: "Total requests",
})