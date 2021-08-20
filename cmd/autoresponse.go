/* autoresponse.go
** Pain Bot
** Author: Jake Sigman
** This file contains the code for the autoresponse function.
 */

package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/clinet/discordgo-embed"
	"os"
	"strings"
)

// Primary function

func Autoresponse(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string) {

	// Delete invocation

	DeleteInvocation(s, m)

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

	if _, err := f.Write([]byte("\n" + write)); err != nil {
		fmt.Println("Error writing to file!\n" + err.Error())
		return
	}
	// Embed construction

	description := "**Prompt:** " + args[1] + "\n**Response:** " + autoresponses[args[1]]

	autoresponse_embed := embed.NewEmbed().
		SetTitle("Autoresponse Added!").
		SetDescription(description).
		SetColor(3066993).
		SetThumbnail("https://github.com/jakesig/Pain-Bot/blob/master/share/icon.png?raw=true").
		MessageEmbed

		// Embed sending

	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, autoresponse_embed); err != nil {
		fmt.Println("Error sending Embed!\n" + err.Error())
		return
	}

	// Close the init.txt file

	if err := f.Close(); err != nil {
		fmt.Println("Error closing file!\n" + err.Error())
		return
	}
}
