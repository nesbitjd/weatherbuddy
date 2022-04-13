package discord

import (
	"fmt"
	"weatherbuddy/internal/weather"

	"github.com/bwmarrin/discordgo"
)

// WeatherData is the struct for the weather request data
type WeatherData struct {
	City string
	Temp string
}

// KtoC defines the conversion between Kelvin and Celsius
const KtoC = -273

// HandleWeather handles the weather request
func EmbedWeather(results weather.Results) *discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{}

	temp := fmt.Sprintf("%f", results.Main.Temp+KtoC)

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   results.Name,
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
