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
