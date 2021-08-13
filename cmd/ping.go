/* ping.go
** Go Bot
** Author: Jake Sigman
** This file contains the code for pinging the bot.
 */

package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Primary function

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Send message replying to user with pong

	if _, err := s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+">, "+"pong!"); err != nil {
		fmt.Println("Error sending message!\n" + err.Error())
	}
}
