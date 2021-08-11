package cmd

import (
  "strings"
  "github.com/bwmarrin/discordgo"
)

func Cmd(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string) {
  switch(strings.Split(m.Content, " ")[0]) {
    case "!ping":
      s.ChannelMessageSend(m.ChannelID, "<@" + m.Author.ID + ">, " + "pong!")
      return
    case "!autoresponse":
      Autoresponse(s, m, autoresponses)
      return
    default:
      break
  }
}