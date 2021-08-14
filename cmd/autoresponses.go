package cmd

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/clinet/discordgo-embed"
)

// Primary function

func Autoresponses(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string) {

	// Delete invocation

	DeleteInvocation(s, m)

	// Start typing

	if err := s.ChannelTyping(m.ChannelID); err != nil {
		fmt.Println("Error starting channel typing!")
		return
	}

	// Writing embed description

	description := ""

	for key, value := range autoresponses {
		description += "**Prompt: **" + key + "\n**Response: **" + value + "\n\n"
	}

	// Embed construction

	autoresponse_embed := embed.NewEmbed().
		SetTitle("Pain Bot - List of Autoresponses").
		SetDescription(description).
		SetColor(15844367).
		SetThumbnail("https://github.com/jakesig/Pain-Bot/blob/master/share/icon.png?raw=true").
		MessageEmbed

		// Embed sending

	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, autoresponse_embed); err != nil {
		fmt.Println("Error in sending embed!\n" + err.Error())
	}

}
