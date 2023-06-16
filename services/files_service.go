package services

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

var fileName = "listUsers.json"

func ReadFileData() map[string][]string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return map[string][]string{}
	}

	var jsonList map[string][]string
	err = json.Unmarshal(file, &jsonList)

	return jsonList
}

func WriteFileData(file *os.File, data []byte) bool {
	if file == nil {
		var err error
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		err = file.Truncate(0)
		_, err = file.Seek(0, 0)

		if err != nil {
			log.Fatal(err)
		}
	}

	_, err := file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	return true
}

func CreateFile() *os.File {
	file, err := os.Create("listUsers.json")
	if err != nil {
		log.Fatal(err)
	}

	return file
}
