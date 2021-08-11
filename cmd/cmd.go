/* cmd.go
** Go Bot
** Author: Jake Sigman
** This file contains the code for processing commands.
 */

package cmd

// Imports

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Primary function

func Cmd(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string) {

	// Switch statement for processing commands

	switch strings.Split(m.Content, " ")[0] {
	case "!ping":
		s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+">, "+"pong!")
		return
	case "!autoresponse":
		Autoresponse(s, m, autoresponses)
		return
	default:
		break
	}
}
