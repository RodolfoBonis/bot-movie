package utils

import (
	"regexp"
	"strings"
)

func NormalizeAndLowercase(s string) string {
	normalized := removeAccents(s)

	return strings.ToLower(normalized)
}

func removeAccents(text string) string {
	accents := map[string]string{
		"á": "a",
		"à": "a",
		"â": "a",
		"ã": "a",
		"é": "e",
		"è": "e",
		"ê": "e",
		"í": "i",
		"ì": "i",
		"î": "i",
		"ó": "o",
		"ò": "o",
		"ô": "o",
		"õ": "o",
		"ú": "u",
		"ù": "u",
		"û": "u",
		"ç": "c",
	}

	pattern := "[" + strings.Join(getKeys(accents), "") + "]"

	regex := regexp.MustCompile(pattern)

	textWithoutAccents := regex.ReplaceAllStringFunc(text, func(match string) string {
		return accents[match]
	})

	return textWithoutAccents
}

func getKeys(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
