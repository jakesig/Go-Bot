package config

import (
	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

  //Checks if the bot sent the message

	if m.Author.ID == s.State.User.ID {
		return
	}

  switch m.Content {
    case "ping":
      s.ChannelMessageSend(m.ChannelID, "Pong!")
    case "pong":
      s.ChannelMessageSend(m.ChannelID, "Ping!")
    default:
      break
  }

}