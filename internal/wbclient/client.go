package wbclient

import (
	"fmt"
	"os"
	"strings"
	"weatherbuddy/internal/discord"
	"weatherbuddy/internal/weather"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Client struct {
	Config weather.Config
	DGo    struct {
		Bot_Token string
		Dg        *discordgo.Session
	}
}

// ParseConfig fills the config struct with the api and discordgo information
func (c *Client) ParseConfig() error {

	var err error
	var exists bool

	c.DGo.Bot_Token, exists = os.LookupEnv("BOT_TOKEN")
	if !exists {
		return fmt.Errorf("BOT_TOKEN not set")
	}

	// Create a new discord session with weatherbuddy
	c.DGo.Dg, err = discordgo.New("Bot " + c.DGo.Bot_Token)
	if err != nil {
		return fmt.Errorf("failed to create discord session: %w", err)
	}

	api_key, exists := os.LookupEnv("API_KEY")
	if !exists {
		return fmt.Errorf("API_KEY not set")
	}

	c.Config = weather.Config(api_key)

	return nil
}

func (c *Client) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Make m.Content lowercase
	message := strings.SplitAfterN(strings.ToLower(m.Content), " ", 2)

	// If the message is "brawl" reply with heroes
	if message[0] == "!weather " {
		logrus.Infof("request recieved, city requested: %s", message[1])
		result, err := c.HandleWeather(message)
		if err != nil {
			logrus.Errorf("unable to handle weather: %w", err)
		}

		body := discord.EmbedWeather(result)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, body)
		if err != nil {
			logrus.Errorf("unable to send embed: %w", err)
		}
	}

}

func (c *Client) HandleWeather(m []string) (weather.Results, error) {

	var results weather.Results

	input, err := handleWeatherInput(m)
	if err != nil {
		return results, fmt.Errorf("unable to handle weather input: %w", err)
	}

	results, err = weather.GetWeather(input, c.Config)
	if err != nil {
		return results, fmt.Errorf("unable to get weather results: %w", err)
	}

	return results, nil

}

func handleWeatherInput(message []string) (string, error) {

	if len(message) == 1 {
		return "", fmt.Errorf("no city name included in request")
	}

	input := strings.ToLower(message[1])
	return input, nil

}
