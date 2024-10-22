package config

import (
	"encoding/json"
	"fmt"

	// "log"
	"os"
	"path/filepath"
)

// Config structure holds all the configuration from the config.json file
type Config struct {
	DB     DatabaseConfig `json:"db"`
	API    APIConfig      `json:"api"`
	Allure AllureConfig   `json:"allure"`
}

// DatabaseConfig holds the database-specific configuration
type DatabaseConfig struct {
	URL                string `json:"url"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
	MaxIdleConnections int    `json:"maxIdleConnections"`
}

// APIConfig holds the API-specific configuration
type APIConfig struct {
	BaseURL string `json:"base_url"`
	ApiUser string `json:"api_user"`
	ApiPass string `json:"api_pass"`
	Timeout int    `json:"timeout"`
}

// AllureConfig holds Allure-related settings
type AllureConfig struct {
	Enable     bool   `json:"enable"`
	ResultsDir string `json:"results_directory"`
}

var config *Config

// LoadConfig loads the configuration from config.json or environment variables if they exist
func LoadConfig() (*Config, error) {
	if config != nil {
		return config, nil // Return if already loaded
	}

	// Set the default path for config.json
	configPath := filepath.Join(".", "config", "config.json")

	// Open config.json file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not open config.json: %v", err)
	}
	defer file.Close()

	// Decode the JSON file into the config struct
	config = &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, fmt.Errorf("could not decode config.json: %v", err)
	}

	// Override with environment variables if they exist
	if dbURL := os.Getenv("DB_URL"); dbURL != "" {
		config.DB.URL = dbURL
	}
	if apiURL := os.Getenv("API_URL"); apiURL != "" {
		config.API.BaseURL = apiURL
	}
	if apiUser := os.Getenv("API_USER"); apiUser != "" {
		config.API.ApiUser = apiUser
	}
	if apiPass := os.Getenv("API_PASS"); apiPass != "" {
		config.API.ApiPass = apiPass
	}
	if allureDir := os.Getenv("ALLURE_RESULTS_DIRECTORY"); allureDir != "" {
		config.Allure.ResultsDir = allureDir
	}

	return config, nil
}

// GetEnv returns the value from the environment variables or from the config
func GetEnv(key string) string {
	// First, check if the value is loaded in config
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return ""
	}

	// Check for specific keys in the loaded configuration
	switch key {
	case "DB_URL":
		return cfg.DB.URL
	case "API_URL":
		return cfg.API.BaseURL
	case "API_USER":
		return cfg.API.ApiUser
	case "API_PASS":
		return cfg.API.ApiPass
	case "ALLURE_RESULTS_DIRECTORY":
		return cfg.Allure.ResultsDir
	default:
		// If not found in config, return from environment variables or empty string if not set
		value := os.Getenv(key)
		return value
	}
}
