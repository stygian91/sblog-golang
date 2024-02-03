package config

import (
	"os"
	"strings"
)

var (
	UrlBase string
)

func Init() error {
	content, err := readEnvFile()

	if err != nil {
		return err
	}

	values := parseEnvFile(content)
	UrlBase = envOrDefault(values, "URL_BASE", "http://localhost/")

	return nil
}

func envOrDefault(values map[string]string, key string, fallback string) string {
	val, exists := values[key]
	if !exists {
		return fallback
	}

	return val
}

func readEnvFile() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path = path + "/.env"
	content, err := os.ReadFile(path)

	return string(content), err
}

func parseEnvFile(content string) map[string]string {
	values := map[string]string{}
	for _, line := range strings.Split(content, "\n") {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}

		values[parts[0]] = parts[1]
	}

	return values
}
