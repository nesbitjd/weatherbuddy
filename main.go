package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		panic(fmt.Errorf("BOT_TOKEN not set"))
	}

	// Create a new discord session with heroes-bot
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Register the brawl func as a callback for "brawl" events
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
	if message[0] == "!hello" {
		body := "Hi there!"
		s.ChannelMessageSend(m.ChannelID, body)
	}

}
