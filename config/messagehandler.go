/* messagehandler.go
** Go Bot
** Author: Jake Sigman
** This file contains the code for handling messages.
 */

package config

// Imports

import (
	"GoBot/cmd"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Message Handler

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Checks if the bot sent the message

	if m.Author.ID == s.State.User.ID {
		return
	}

	// Process it as a command first

	cmd.Cmd(s, m, autoresponses)

	// For loop for autoresponses

	for key, value := range autoresponses {
		if strings.Contains(m.Content, key) && m.Content[:1] != "!" {
			s.ChannelMessageSend(m.ChannelID, value)
		}
	}
}
