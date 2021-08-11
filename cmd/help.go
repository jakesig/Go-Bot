package cmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/clinet/discordgo-embed"
)

var (
	help_embed *discordgo.MessageEmbed
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Embed construction

	help_embed = embed.NewGenericEmbedAdvanced("Go Bot List of Commands",
		"**$help:** Opens this menu.\n"+
			"**$ping:** Pings the bot.\n"+
			"**$autoresponse `{prompt}` `{response}`:** Adds autoresponse to bot.",
		15844367)
	s.ChannelMessageSendEmbed(m.ChannelID, help_embed)
}
