package cmd

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"main.go/internal"
)

// WeatherData is the struct for the weather request data
type WeatherData struct {
	City string
	Temp string
}

// HandleWeather handles the weather request
func HandleWeather(message []string) *discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{}

	api_key := os.Getenv("API_KEY")
	if api_key == "" {
		embed := &discordgo.MessageEmbed{
			Title:       "Weatherbuddy",
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xad0000,
			Description: "API Key not set",
			//Fields:      fields,
		}
		return embed
	}

	input, err := handleWeatherInput(message)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title:       "Weatherbuddy",
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xad0000,
			Description: err.Error(),
			//Fields:      fields,
		}
		return embed
	}

	weather, err := internal.GetWeather(input)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title:       "Weatherbuddy",
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xad0000,
			Description: err.Error(),
			//Fields:      fields,
		}
		return embed
	}

	temp := fmt.Sprintf("%f", weather.Main.Temp)

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   input,
		Value:  temp,
		Inline: true,
	})

	// Populate embed with fields
	embed := &discordgo.MessageEmbed{
		Title:       "Weatherbuddy",
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x36A8DE,
		Description: "Here's the weather!",
		Fields:      fields,
	}
	return embed
}

func handleWeatherInput(message []string) (string, error) {

	input := message[1]

	return input, nil
}
