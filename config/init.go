/* cmd.go
** Pain Bot
** Author: Jake Sigman
** This file contains the code for initializing the bot.
 */

package config

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"os"
	"strconv"
	"strings"
)

// Variables

var (
	token         string
  ignoreIDs     []string
  ignore        bool
	section       string
	prefix        string
	paincount     int
	linecount     int
	autoresponses map[string]string
)

// Initialization function

func Init() {

	// Open init.txt file

	f, _ := os.Open("./init.txt")

	// Variables to keep track of lines, linecount, and autoresponses

	buf := make([]byte, 1024)
	linecount = 0
	autoresponses = make(map[string]string)
	prefix = "$"

	// Loop for reading init.txt

	for {

		n, err := f.Read(buf)

		if err == io.EOF {
			break
		}

		if n > 0 {

			// Log to the console what we're up to

			fmt.Println("GO: Reading config file...")

			// Get the lines of the file

			lines := strings.Split(string(buf[:n]), "\n")

			// Parse the lines of the file

			for i := 0; i < len(lines); i++ {
				line := lines[linecount]
				linecount = linecount + 1

				// The first line has the token on it

				if linecount == 1 {
					token = strings.Split(line, ": ")[1]
				}

        // Checking servers that we want to ignore for the paincount

        if strings.TrimSpace(line) == "IGNORE" {
					section = "IGNORE"
				} else if section == "IGNORE" {
					ignoreIDs = append(ignoreIDs, strings.TrimSpace(line))
				}

				// Checking if we hit the autoresponses section of the init file

				if strings.TrimSpace(line) == "AUTORESPONSES" {
					section = "AUTORESPONSES"
				} else if section == "AUTORESPONSES" {
					components := strings.Split(strings.TrimSpace(line), "/")
					autoresponses[components[0]] = components[1]
				}
			}
		}
	}

	// Close the file

	if err := f.Close(); err != nil {
		fmt.Println("Error closing init.txt!\n" + err.Error())
		return
	}

	// Open the file keeping track of the pain count, and read the count

	f, _ = os.Open("./count.txt")
	n, _ := f.Read(buf)
	paincount, _ = strconv.Atoi(strings.TrimSpace(string(buf[:n])))

	// Creates a new discordgo client

	dg, err := discordgo.New("Bot " + strings.TrimSpace(token))

	if err != nil {
		fmt.Println("Unsuccessful creation of bot!\n" + err.Error())
		return
	}

	// Handlers

	dg.AddHandler(messageCreate)

	// Intents

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open connection to Discord

	if err := dg.Open(); err != nil {
		fmt.Println("Error opening Bot!\n" + err.Error())
		return
	}

	// Set status of the bot

	if err := (*discordgo.Session).UpdateGameStatus(dg, 0, "$pain"); err != nil {
		fmt.Println("Error setting status!\n" + err.Error())
		return
	}

	// Logs to the console once the bot is running

	fmt.Println("GO: Logged in as " + dg.State.User.Username + "#" + dg.State.User.Discriminator)

	// Run a function asynchronously to send pain in #general once a day

	PainEveryDay(dg)
}
