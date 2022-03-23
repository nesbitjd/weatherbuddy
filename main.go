package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"main.go/cmd"
)

func main() {
	token, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		panic(fmt.Errorf("BOT_TOKEN not set"))
	}

	// Create a new discord session with weatherbuddy
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dg.AddHandler(handler)

	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	dg.Close()
}

func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Make m.Content lowercase
	message := strings.Split(strings.ToLower(m.Content), " ")

	// If the message is "brawl" reply with heroes
	if message[0] == "!weather" {
		body := cmd.HandleWeather(message)
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, body)
		if err != nil {
			logrus.Errorf("unable to send embed: %w", err)
		}
	}

}
