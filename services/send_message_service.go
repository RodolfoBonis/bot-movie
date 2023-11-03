package services

import "fmt"

func GetMessage(jsonList map[string][]string) string {
	if jsonList == nil {
		data := ReadFileData("listUsers.json")
		jsonList = data
	}

	message := fmt.Sprintf("Quem irá escolher o filme hoje será <@%s>", jsonList["users"][0])

	return message
}
