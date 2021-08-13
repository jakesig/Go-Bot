package cmd

// Imports

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

// Primary function

func Pain(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Open the file

	file, err := os.Open("pain.jpg")
	if err != nil {
		fmt.Println("Error opening file!\n" + err.Error())
		return
	}

	// Send the file

	if _, err = s.ChannelFileSend(m.ChannelID, "pain.jpg", file); err != nil {
		fmt.Println("Error sending image!\n" + err.Error())
		return
	}
}
