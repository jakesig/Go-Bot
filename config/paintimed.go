/* paintimed.go
** Pain Bot
** Author: Jake Sigman
** This file contains the code for saying pain once a day at a random time.
 */

package config

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strconv"
	"time"
)

// Variables

var (
	paintoday        bool
	general_channels []string
)

// Primary function

func PainEveryDay(s *discordgo.Session) {

	fmt.Println("GO: Initializing times...")
	paintoday = false

	// Get a random timestamp

	paintime := randomTimestamp()
	hour, min, _ := paintime.Clock()
	fmt.Println("GO: Initial time set: " + strconv.Itoa(hour) + " hours and " + strconv.Itoa(min) + " minutes")

	// Find channels named "general"

	findGeneral(s)

	// Go routine for sending pain once every day at a random time

	go func() {
		for {
			if m, h, s := time.Now().Clock(); m == 0 && h == 0 && s == 0 {
				paintoday = false
				paintime = randomTimestamp()
				hour, min, _ := paintime.Clock()
				fmt.Println("GO: Time set: " + strconv.Itoa(hour) + " hours and " + strconv.Itoa(min) + " minutes")
			}

			nhour, nmin, _ := time.Now().Clock()
			hour, min, _ := paintime.Clock()

			for i := 0; i < len(general_channels); i++ {
				if nhour == hour && nmin == min && !paintoday {
					if _, err := s.ChannelMessageSend(general_channels[i], "pain"); err != nil {
						fmt.Println("Error sending message!\n" + err.Error())
						return
					}
					paintoday = true
				}
			}

		}
	}()
}

// Function that generates a random timestamp

func randomTimestamp() time.Time {
	rand.Seed(time.Now().UnixNano())
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	return randomNow
}

// Function for determining all channels named "general"

func findGeneral(s *discordgo.Session) []string {
	for _, guild := range s.State.Guilds {
		channels, _ := s.GuildChannels(guild.ID)
		for _, c := range channels {
			if c.Name == "general" {
				general_channels = append(general_channels, c.ID)
			}
		}
	}

	return general_channels
}
