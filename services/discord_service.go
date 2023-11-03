package services

import (
	"encoding/json"
	"fmt"
	"github.com/RodolfoBonis/bot_movie/config"
	"github.com/RodolfoBonis/bot_movie/entities"
	"github.com/RodolfoBonis/bot_movie/utils"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"os"
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
		{
			Name:        "search-course",
			Description: "This Command send a channel message with the course that you want to search",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "search",
					Description: "Search the course that you want",
					Required:    true,
				},
			},
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
			options := i.ApplicationCommandData().Options

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			margs := make([]interface{}, 0, len(options))

			if option, ok := optionMap["search"]; ok {
				margs = append(margs, option.StringValue())
			}

			go func(s *discordgo.Session, i *discordgo.InteractionCreate) {

				userChannel, _ := s.UserChannelCreate(i.Member.User.ID)

				courses := getCoursesList(margs[0].(string))

				for _, item := range courses {
					output := fmt.Sprintf(" %s", item.Name)
					links := ""
					for _, link := range item.Links {
						links += fmt.Sprintf("%s: %s\n", link.Name, link.Link)
					}

					embed := &discordgo.MessageEmbed{
						Title:       output,
						Description: links,
					}

					_, err := s.ChannelMessageSendEmbed(userChannel.ID, embed)
					if err != nil {
						log.Fatal(err)
						return
					}
				}

			}(s, i)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Lhe enviamos um DM com o resultado da sua pesquisa",
				},
			})

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

func getCoursesList(searchTerm string) []entities.CoursesData {
	file, err := os.Open("courses.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo JSON:", err)
		return nil
	}
	defer file.Close()

	var data []entities.CoursesData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return nil
	}

	var filteredList []entities.CoursesData

	searchTerm = utils.NormalizeAndLowercase(searchTerm)

	for _, item := range data {
		normalizedTerm := utils.NormalizeAndLowercase(item.Name)
		if strings.Contains(normalizedTerm, searchTerm) {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}

func SendMessage(message string) {
	_, err := Session.ChannelMessageSend(config.EnvChannelID(), message)

	if err != nil {
		log.Fatal(err)
	}
}
