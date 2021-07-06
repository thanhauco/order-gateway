package main
import "flag"
var version = "1.0.0"
func main() {
    v := flag.Bool("v", false, "version")
    flag.Parse()
    if *v { println(version); return }
    // ... rest
}