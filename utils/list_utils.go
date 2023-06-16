package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func ReorderList(data map[string][]string) []byte {
	array, firstElement := shift(data["users"])

	data["users"] = append(array, firstElement)

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}

func shift(slice []string) ([]string, string) {
	if len(slice) == 0 {
		return slice, "0"
	}
	firstElement := slice[0]
	slice = slice[1:]
	return slice, firstElement
}
