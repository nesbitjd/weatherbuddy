package cmd

import "github.com/bwmarrin/discordgo"

func handleWeather(message []string) *discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{}

	// Check for errors on message input from handleBrawlInput
	input, err := handleWeatherInput(message)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title:       "Brawl Hero Selector",
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xad0000,
			Description: err.Error(),
			//Fields:      fields,
		}
		return embed
	}
}
