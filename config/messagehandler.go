package config

// Imports

import (
  "strings"
	"github.com/bwmarrin/discordgo"
  "GoBot/cmd"
)

// Message Handler

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

  //Checks if the bot sent the message

	if m.Author.ID == s.State.User.ID {
		return
	}

  cmd.Cmd(s, m, autoresponses)

  // For loop for autoresponses

  for key, value := range autoresponses {
    if (strings.Contains(m.Content, key) && m.Content[:1] != "!") {
      s.ChannelMessageSend(m.ChannelID, value)
    }
  }

}