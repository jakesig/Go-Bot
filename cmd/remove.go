/* remove.go
** Pain Bot
** Author: Jake Sigman
** This file contains the code for removing autoresponses.
 */

package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/clinet/discordgo-embed"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Variables

var (
	linecount int
	write     string
	section   string
)

// Primary function

func Remove(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string) {

	// Delete invocation

	DeleteInvocation(s, m)

	// Reading each argument of command

	args := strings.Split(m.Content, " ")

	// Open the init.txt file for saving the autoresponse

	f, _ := os.Open("./init.txt")

	// Embed construction

	description := "**Prompt:** " + args[1] + "\n**Response:** " + autoresponses[args[1]]

	autoresponse_embed := embed.NewEmbed().
		SetTitle("Autoresponse Removed!").
		SetDescription(description).
		SetColor(15158332).
		SetThumbnail("https://github.com/jakesig/Pain-Bot/blob/master/share/icon.png?raw=true").
		MessageEmbed

	// Remove autoresponse from map

	delete(autoresponses, args[1])

	// Embed sending

	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, autoresponse_embed); err != nil {
		fmt.Println("Error sending Embed!\n" + err.Error())
		return
	}

	// Read init.txt

	buf := make([]byte, 1024)
	linecount = 0

	for {

		n, err := f.Read(buf)

		if err == io.EOF {
			break
		}

		if n > 0 {

			// Get the lines of the file

			lines := strings.Split(string(buf[:n]), "\n")

			// Parse the lines of the file

			for i := 0; i < len(lines); i++ {
				line := lines[linecount]
				linecount = linecount + 1

				// Look for the autoresponse to remove

				if strings.TrimSpace(line) == "AUTORESPONSES" {
					section = "AUTORESPONSES"
				} else if section == "AUTORESPONSES" && (args[1] == strings.Split(line, "/")[0] || line == "") {
					continue
				}

				write = write + line + "\n"
			}
		}
	}

	// Overwrite init.txt with updated autoresponses

	if err := ioutil.WriteFile("init.txt", []byte(write), 0644); err != nil {
		fmt.Println("Error writing to file!\n" + err.Error())
		return
	}

	// Close the init.txt file

	if err := f.Close(); err != nil {
		fmt.Println("Error closing file!\n" + err.Error())
		return
	}
}
