package envstr

import (
	"os"
	"regexp"
	"strings"
)

func Apply(s string) (string, error) {
	return apply(s, func(k string) string {
		return os.Getenv(k)
	})
}

func apply(s string, f func(string) string) (string, error) {
	re, err := regexp.Compile(`\${[A-Z][a-zA-Z0-9_]*}`)
	if err != nil {
		return "", err
	}
	result := re.ReplaceAllStringFunc(s, func(match string) string {
		k := strings.TrimSpace(strings.TrimRight(strings.TrimLeft(match, "${"), "$}"))
		return f(k)
	})
	return result, nil
}
