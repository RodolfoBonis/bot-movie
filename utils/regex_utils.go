package utils

import (
	"golang.org/x/text/unicode/norm"
	"strings"
)

func NormalizeAndLowercase(s string) string {
	normalized := norm.NFD.String(s)
	return strings.ToLower(normalized)
}
