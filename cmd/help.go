package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/clinet/discordgo-embed"
)

// Variables

var (
	help_embed *discordgo.MessageEmbed
)

// Primary function

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Embed construction

	help_embed = embed.NewGenericEmbedAdvanced("Go Bot List of Commands",
		"**$help:** Opens this menu.\n"+
			"**$ping:** Pings the bot.\n"+
			"**$autoresponse `{prompt}` `{response}`:** Adds autoresponse to bot.",
		15844367)

	// Embed sending

	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, help_embed); err != nil {
		fmt.Println("Error in sending embed!\n" + err.Error())
	}
}
