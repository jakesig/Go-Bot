/* cmd.go
** Go Bot
** Author: Jake Sigman
** This file contains the code for initializing the bot.
 */

package config

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"os"
	"strings"
)

// Variables

var (
	token         string
	section       string
	linecount     int
	autoresponses map[string]string
)

// Initialization function

func Init() {

	// Open init.txt file

	f, err := os.Open("./init.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	// Variables to keep track of lines, linecount, and autoresponses

	buf := make([]byte, 1024)
	linecount = 0
	autoresponses = make(map[string]string)

	// Loop for reading init.txt

	for {

		n, err := f.Read(buf)

		if err == io.EOF {
			break
		}

		if n > 0 {
			lines := strings.Split(string(buf[:n]), "\n")

			// Get the line of the file

			for i := 0; i < len(lines); i++ {
				line := lines[linecount]
				linecount = linecount + 1

				// The first line has the token on it

				if linecount == 1 {
					token = strings.Split(line, ": ")[1]
				}

				// Checking if we hit the autoresponse section of the init file

				if line == strings.TrimSpace("AUTORESPONSES") {
					section = line
				} else if section == "AUTORESPONSES" {
					components := strings.Split(line, "/")
					autoresponses[components[0]] = components[1]
				}
			}
		}
	}

	// Close the file

	f.Close()

	// Creates a new discordgo client

	dg, err := discordgo.New("Bot " + strings.TrimSpace(token))

	if err != nil {
		fmt.Println("Unsuccessful: ", err)
	}

	// Handlers

	dg.AddHandler(messageCreate)

	// Intents

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open connection to Discord

	dg.Open()

	// Logs to the console once the bot is running

	fmt.Println("Logged in as " + dg.State.User.Username + "#" + dg.State.User.Discriminator)
}
