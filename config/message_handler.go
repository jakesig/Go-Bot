/* message_handler.go
** Pain Bot
** Author: Jake Sigman
** This file contains the code for handling messages.
 */

package config

// Imports

import (
	"PainBot/cmd"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

// Message Handler

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Checks if the bot sent the message

	if m.Author.ID == s.State.User.ID {
		return
	}

	// Checks if message is a DM

	if isDM, _ := ComesFromDM(s, m); isDM {
		dmChannel, _ := s.UserChannelCreate(m.Author.ID)

		// Sends a random message from messages.txt

		if _, err := s.ChannelMessageSend(dmChannel.ID, messages[rand.Intn(len(messages))]); err != nil {
			fmt.Println("Error sending message!\n" + err.Error())
		}

		return
	}

	// Process it as a command first

	cmd.Cmd(s, m, autoresponses, "$", paincount)

	// Determine if the server needs to be ignored for the pain counter

	ignore = false

	for i := 0; i < len(ignoreIDs); i++ {
		if m.GuildID == ignoreIDs[i] {
			ignore = true
			break
		}
	}

	// Pain counter

	if strings.ToLower(m.Content) == "pain" && !ignore {
		paincount++
		if err := ioutil.WriteFile("count.txt", []byte(strconv.Itoa(paincount)), 0644); err != nil {
			fmt.Println("Error writing to file!\n" + err.Error())
			return
		}
	}

	// For loop for autoresponses

	for key, value := range autoresponses {
		if strings.Contains(strings.ToLower(m.Content), key) && m.Content[:1] != prefix {
			if _, err := s.ChannelMessageSend(m.ChannelID, value); err != nil {
				fmt.Println("Error sending message!\n" + err.Error())
				return
			}
		}
	}
}

// ComesFromDM returns true if a message comes from a DM channel
// Source: https://github.com/bwmarrin/discordgo/wiki/FAQ#checking-if-a-message-is-a-direct-message-dm

func ComesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}
