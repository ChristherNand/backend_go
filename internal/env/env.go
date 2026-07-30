package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		valueInt, err := strconv.Atoi(value)
		if err == nil {
			return valueInt
		}
	}
	return fallback
}
