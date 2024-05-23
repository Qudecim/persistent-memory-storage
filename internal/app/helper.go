package app

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func escapeString(str string) string {
	escapedStr := ""
	for _, char := range str {
		if char == '\n' {
			escapedStr += "\\n"
		} else {
			escapedStr += string(char)
		}
	}

	return escapedStr
}

func deescapeString(str string) string {
	return strings.ReplaceAll(str, "\\n", "\n")
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func randomInThouthand(i int) bool {
	return rand.Intn(1000) < i
}

func Timestamp() string {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()
	return strconv.FormatInt(unixTimestamp, 10)
}
