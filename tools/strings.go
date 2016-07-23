package tools

import (
	"strings"
)

func Trim(bt []byte) string {
	return strings.Trim(string(bt), " ")
}
