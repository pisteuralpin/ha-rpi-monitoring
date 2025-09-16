package env

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func LoadEnv() {
	// Load .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		file, err := os.Open(".env")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		sc := bufio.NewScanner(file)

		// Parsing .env file line by line and setting environment variables
		for sc.Scan() {
			line := sc.Text()
			if len(line) == 0 || line[0] == '\n' || line[0] == ' ' || line[0] == '#' {
				continue
			}
			parts := strings.SplitN(strings.SplitAfterN(line, "#", 1)[0], "=", 2)
			if len(parts) == 2 {
				os.Setenv(parts[0], strings.TrimSpace(parts[1]))
			}
		}

		if err := sc.Err(); err != nil {
			panic(err)
		}

	}
}

// getEnv reads an environment variable or returns a default value if not set
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt reads an environment variable as an integer or returns a default value if not set or invalid
func GetEnvAsInt(name string, defaultValue int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// getEnvAsBool reads an environment variable as a boolean or returns a default value if not set or invalid
func GetEnvAsBool(name string, defaultValue bool) bool {
	valueStr := GetEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}
