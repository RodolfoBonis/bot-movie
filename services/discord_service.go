package services

import (
	"fmt"
	"github.com/RodolfoBonis/bot_movie/config"
	"github.com/RodolfoBonis/bot_movie/utils"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"strings"
)

var Session *discordgo.Session

func StartDiscordConnection() {
	var err error

	Session, err = discordgo.New("Bot " + config.EnvBotToken())

	if err != nil {
		log.Fatal(err)
	}

	Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})

	err = Session.Open()

	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	startCommandsHandler()
}

func startCommandsHandler() {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "selector-movie",
			Description: "This Command send a channel message with whom you will choose the movie",
		},
	}

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"selector-movie": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := ReadFileData("listUsers.json")
			message := GetMessage(data)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: message,
				},
			})
			reorderedData := utils.ReorderList(data)

			result := WriteFileData("listUsers.json", nil, reorderedData)
			if result {
				fmt.Println("List reordered with success")
			}
		},
		"search-course": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			args := strings.Fields(i.Message.Content)

			fmt.Println(args)
		},
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := Session.ApplicationCommandCreate(Session.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func getCoursesList() {
	//file, err := os.Open("courses.json")
	//if err != nil {
	//	fmt.Println("Erro ao abrir o arquivo JSON:", err)
	//	return
	//}
	//defer file.Close()
	//
	//var data []entities.CoursesData
	//decoder := json.NewDecoder(file)
	//if err := decoder.Decode(&data); err != nil {
	//	fmt.Println("Erro ao decodificar o JSON:", err)
	//	return
	//}
	//
	//searchIndex := make(map[string][]entities.CoursesData)
}

func SendMessage(message string) {
	_, err := Session.ChannelMessageSend(config.EnvChannelID(), message)

	if err != nil {
		log.Fatal(err)
	}
}
