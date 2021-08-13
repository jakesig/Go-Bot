package cmd

import (
  "fmt"
  "github.com/bwmarrin/discordgo"
  "github.com/clinet/discordgo-embed"
)

var (
  description string
)

func Autoresponses(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string) {

  // Delete invocation

	DeleteInvocation(s, m)

  // Start typing

  s.ChannelTyping(m.ChannelID)

  // Writing embed description

  description = ""

  for key, value := range autoresponses {
    description += "**Prompt: **" + key + "\n**Response: **" + value + "\n\n"
	}

  // Embed construction

	help_embed = embed.NewGenericEmbedAdvanced("Pain Bot - List of Autoresponses", description, 15844367)

  // Embed sending

	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, help_embed); err != nil {
		fmt.Println("Error in sending embed!\n" + err.Error())
	}

}