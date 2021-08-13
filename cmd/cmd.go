/* cmd.go
** Go Bot
** Author: Jake Sigman
** This file contains the code for processing commands.
 */

package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Primary function

func Cmd(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string, prefix string, paincount int) {

	// Switch statement for processing commands

	switch strings.Split(m.Content, " ")[0] {
	case prefix + "ping":
		Ping(s, m)
		break
	case prefix + "autoresponse":
		Autoresponse(s, m, autoresponses)
		break
	case prefix + "help":
		Help(s, m)
		break
	case prefix + "paincount":
		PainCount(s, m, paincount)
		break
	case prefix + "pain":
		Pain(s, m)
		break
	default:
		break
	}
}

// Function for deleting invocations

func DeleteInvocation(s *discordgo.Session, m *discordgo.MessageCreate) {
	if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
		fmt.Println("Error deleting message!\n" + err.Error())
		return
	}
}
