/* autoresponse.go
** Go Bot
** Author: Jake Sigman
** This file contains the code for the autoresponse function.
 */

package cmd

// Imports

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

var (
	embed *discordgo.MessageEmbed
)

// Primary function

func Autoresponse(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string) {

	// Reading each argument of command, taking the first argument as the key, and second as the value in the autoresponses map

	args := strings.Split(m.Content, " ")
	write := args[1] + "/" + args[2]
	autoresponses[args[1]] = args[2]

	// Open the init.txt file for saving the autoresponse

	f, _ := os.OpenFile("init.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// Parses any other arguments, adding them on to the value

	if len(args) > 3 {
		for i := 3; i < len(args); i++ {
			write += " " + args[i]
		}
		key := strings.Split(write, "/")
		autoresponses[args[1]] = key[1]
	}

	// Saving key and value to init.txt

	f.Write([]byte("\n" + write))

	//embed = discordgo.MessageEmbed{}
	//embed.Title = "Autoresponse Added!"
	//embed.Type = discordgo.EmbedTypeRich
	//embed.Description = "**Prompt:** " + args[1] + "\n**Response:** " + autoresponses[args[1]]
	//embed.Timestamp = time.Now().String()
	//
	//s.ChannelMessageSendEmbed(m.ChannelID, embed)

	// Close the init.txt file

	f.Close()
}
