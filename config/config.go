package config
import "errors"
func (c *Config) Validate() error {
    if c.Port == 0 { return errors.New("invalid port") }
    return nil
}