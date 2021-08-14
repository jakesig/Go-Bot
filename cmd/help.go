/* help.go
** Pain Bot
** Author: Jake Sigman
** This file contains the code for the help menu.
 */

package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/clinet/discordgo-embed"
)

// Primary function

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Delete invocation

	DeleteInvocation(s, m)

	// Start typing

	if err := s.ChannelTyping(m.ChannelID); err != nil {
		fmt.Println("Error starting channel typing!")
		return
	}

	// Embed construction

	description := "**$help:** Opens this menu.\n" +
		"**$ping:** Pings the bot.\n" +
		"**$paincount:** Informs user how many times \"pain\" was said.\n" +
		"**$pain:** Pain.\n" +
		"**$autoresponse `{prompt}` `{response}`:** Adds autoresponse to bot.\n" +
		"**$autoresponses**: Sends list of all current autoresponses."

	help_embed := embed.NewEmbed().
		SetTitle("Pain Bot - List of Commands").
		SetDescription(description).
		SetColor(15844367).
		SetThumbnail("https://github.com/jakesig/Pain-Bot/blob/master/share/icon.png?raw=true").
		MessageEmbed

	// Embed sending

	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, help_embed); err != nil {
		fmt.Println("Error in sending embed!\n" + err.Error())
	}
}
