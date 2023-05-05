package envstr

import (
	"fmt"
	"os"
	"strings"
)

func Apply(s string) string {
	var found bool
	for {
		s, found = ApplyOnce(s)
		if !found {
			return s
		}
	}
}

func ApplyOnce(s string) (string, bool) {
	idxS := strings.Index(s, "${")
	if idxS == -1 {
		return s, false
	}
	// 開始位置以降に限定
	idxEFromS := strings.Index(s[idxS+len("${"):], "}")
	if idxEFromS == -1 {
		return s, false
	}
	// 開始位置以降の相対距離に、開始位置の距離を補正
	idxE := idxS + idxEFromS + len("}")
	envK := s[idxS+len("${") : idxE+len("}")]
	envV := os.Getenv(envK)
	fmt.Println(envK, envV) // TODO
	return s[:idxS] + envV + s[idxE+len("}")+1:], true
}
