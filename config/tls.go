package config
import "crypto/tls"
func LoadTLS() *tls.Config {
    // Load certs
    return &tls.Config{MinVersion: tls.VersionTLS13}
}