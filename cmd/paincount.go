package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/clinet/discordgo-embed"
	"strconv"
)

// Variables

var (
	paincount_embed *discordgo.MessageEmbed
)

// Primary function

func PainCount(s *discordgo.Session, m *discordgo.MessageCreate, paincount int) {

	// Embed construction

	paincount_embed = embed.NewGenericEmbedAdvanced("Pain Count", strconv.Itoa(paincount), 15844367)

	// Embed sending

	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, paincount_embed); err != nil {
		fmt.Println("Error sending embed!\n" + err.Error())
		return
	}
}
