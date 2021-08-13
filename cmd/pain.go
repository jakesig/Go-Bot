/* pain.go
** Go Bot
** Author: Jake Sigman
** Pain.
 */

package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Primary function

func Pain(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Delete invocation

	DeleteInvocation(s, m)

	// Send the image

	if _, err := s.ChannelMessageSend(m.ChannelID, "https://github.com/jakesig/Pain-Bot/blob/master/share/pain.jpg?raw=true"); err != nil {
		fmt.Println("Error sending image!\n" + err.Error())
		return
	}
}
