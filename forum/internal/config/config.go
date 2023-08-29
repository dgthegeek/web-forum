package config
import (
	"log"
	"os"
	"strconv"
)
type Config struct {
	Port         int
	DatabasePath string
	// Add more configuration options as needed
}
func NewConfig() *Config {
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Invalid PORT value")
	}
	return &Config{
		Port:         port,
		DatabasePath: getEnv("DATABASE_PATH", "database.db"),
		// Initialize other configuration options
	}
}
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}