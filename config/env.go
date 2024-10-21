package config

import "os"

// GetEnv returns the value of the environment variable if it exists, otherwise it returns a default value.
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
