/* bot_join_handler.go
** Pain Bot
** Author: Jake Sigman
** This file contains the code for sending a message when the bot joins.
 */

package config

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/clinet/discordgo-embed"
	"time"
)

var (
	generalID string
)

// Message Handler

func guildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {

	// Determine if the bot was in the server already

	t, _ := discordgo.Timestamp.Parse(g.JoinedAt)
	hour, min, _ := t.Clock()
	nhour, nmin, _ := time.Now().Clock()

	if hour != nhour && min != nmin {
		return
	}

	// Determine which channel is called general

	channels, _ := s.GuildChannels(g.ID)
	for _, c := range channels {
		if c.Name == "general" {
			generalID = c.ID
		}
	}

	// Embed construction

	description := "By adding this bot you have added pain to your server.\n\n" +
		"__**Features**__\n" +
		"**Pain every day:** I'll send \"pain\" at a random time in the day.\n" +
		"**Instant pain:** Every time somebody sends a message containing \"pain\", I'll respond \"agony, even\".\n" +
		"**Pain on-demand:** Type $pain for on-demand pain.\n\n" +
		"Type $help for more info!"

	joinEmbed := embed.NewEmbed().
		SetTitle("Pain. Agony, even.").
		SetDescription(description).
		SetColor(15844367).
		SetThumbnail("https://github.com/jakesig/Pain-Bot/blob/master/share/icon.png?raw=true").
		MessageEmbed

	// Embed sending

	if _, err := s.ChannelMessageSendEmbed(generalID, joinEmbed); err != nil {
		fmt.Println("Error in sending embed!\n" + err.Error())
	}
}
