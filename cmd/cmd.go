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

func Cmd(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string, prefix string) {

	// Switch statement for processing commands

	switch strings.Split(m.Content, " ")[0] {
	case prefix + "ping":
		s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+">, "+"pong!")
		break
	case prefix + "autoresponse":
		Autoresponse(s, m, autoresponses)
		break
	case prefix + "help":
		Help(s, m)
		break
	default:
		break
	}
}
