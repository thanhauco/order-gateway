package config
import "github.com/spf13/viper"
func LoadViper() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.ReadInConfig()
}