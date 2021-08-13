/* paincount.go
** Pain Bot
** Author: Jake Sigman
** This file contains the code for the pain count function.
 */

package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/clinet/discordgo-embed"
	"strconv"
)

// Primary function

func PainCount(s *discordgo.Session, m *discordgo.MessageCreate, paincount int) {

	// Delete invocation

	DeleteInvocation(s, m)

	// Embed construction

  paincount_embed := embed.NewEmbed().
    SetTitle("Pain Counter").
    SetDescription(strconv.Itoa(paincount)).
    SetColor(15844367).
    SetThumbnail("https://github.com/jakesig/Pain-Bot/blob/master/share/icon.png?raw=true").
    MessageEmbed

	// Embed sending

	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, paincount_embed); err != nil {
		fmt.Println("Error sending embed!\n" + err.Error())
		return
	}
}
