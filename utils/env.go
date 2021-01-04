package utils

import "os"

// GetEnv : get environment variables with default value
func GetEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

// RequireGetEnv : get environment variable with requirement
func RequireGetEnv(key string) (val string, exist bool) {
	val, exist = os.Getenv(key), true
	if val == "" {
		exist = false
	}
	return
}

// DefaultStringEmpty : get default string if value is empty
func DefaultStringEmpty(value string, defaultVal string) string {
	if value == "" {
		value = defaultVal
	}
	return value
}

// SetIfNotEmpty : set target = value if value is not empty
func SetIfNotEmpty(target *string, value string) {
	if value != "" {
		*target = value
	}
}
