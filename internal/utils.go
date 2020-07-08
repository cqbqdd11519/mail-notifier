package internal

import (
	"fmt"
	"os"
	"strings"
)

const (
	EnvUserPrefix string = "appr-"
)

func CheckEnv(keys []string) error {
	for _, k := range keys {
		if os.Getenv(k) == "" {
			return fmt.Errorf("environment variable %s should be set", k)
		}
	}

	return nil
}

func ParseUserEnv() map[string]string {
	result := map[string]string{}

	prefixLen := len(EnvUserPrefix)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) < 1 {
			continue
		}
		if strings.HasPrefix(pair[0], EnvUserPrefix) {
			result[pair[0][prefixLen:]] = pair[1]
		}
	}

	return result
}
