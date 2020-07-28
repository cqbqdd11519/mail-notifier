package internal

import (
	"fmt"
	"os"
)

func CheckEnv(keys []string) error {
	for _, k := range keys {
		if os.Getenv(k) == "" {
			return fmt.Errorf("environment variable %s should be set", k)
		}
	}

	return nil
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
