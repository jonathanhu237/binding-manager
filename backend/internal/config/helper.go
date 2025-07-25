package config

import (
	"fmt"
	"os"
	"strconv"
)

func getStringEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("Environment variable %s is not set", key)
	}
	return value, nil
}

func getIntEnv(key string) (int, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return 0, fmt.Errorf("Environment variable %s is not set", key)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("Environment variable %s is not a valid integer: %v", key, err)
	}
	return intValue, nil
}
