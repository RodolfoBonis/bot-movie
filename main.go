package main

import (
	"encoding/json"
	"github.com/RodolfoBonis/bot_movie/config"
	"github.com/RodolfoBonis/bot_movie/services"
	"log"
	"math/rand"
	"os"
	"os/signal"
)

func init() {
	config.LoadEnvVars()
	saveFirstTimeData()
	services.StartDiscordConnection()
}

func main() {
	services.ScheduleBotRoutine()

	defer services.Session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

}

func saveFirstTimeData() {
	userList := []string{
		config.EnvAlnWolfID(),
		config.EnvErmaclessID(),
		config.EnvLaisID(),
		config.EnvJonasID(),
	}

	rand.Shuffle(len(userList), func(i, j int) {
		userList[i], userList[j] = userList[j], userList[i]
	})

	data := map[string][]string{
		"users": userList,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	file := services.CreateFile()
	services.WriteFileData(file, jsonData)
}
