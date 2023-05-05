package envstr

import (
	"os"
	"regexp"
	"strings"
)

func Apply(s string) string {
	re, _ := regexp.Compile(`\${[A-Z][a-zA-Z0-9_]*}`)
	result := re.ReplaceAllStringFunc(s, func(match string) string {
		k := strings.TrimSpace(strings.TrimRight(strings.TrimLeft(match, "${"), "$}"))
		v := os.Getenv(k)
		return v
	})
	return result
}
