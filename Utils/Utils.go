package Utils

import "os"

func GetEnvDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		value = defaultValue
	}
	return value
}
