package config

import (
  "io"
  "os"
  "strings"
  "log"
  "fmt"
	"github.com/bwmarrin/discordgo"
)

var token string

func Init() {
  f, err := os.Open("../init.txt")

  if err != nil {
      log.Fatalf("unable to read file: %v", err)
  }

  defer f.Close()

  buf := make([]byte, 1024)

  for {
    n, err := f.Read(buf)

    if err == io.EOF {
      break
    }

    if n > 0 {
      line := strings.Split(string(buf[:n]), ": ")

      if line[0] == "TOKEN" {
        token = line[1];
      }
    }

  }

  // Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)

  if err != nil {
    fmt.Println("Unsuccessful: ", err)
  }

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	dg.Open()

	fmt.Println("Logged in as " + dg.State.User.Username + "#" + dg.State.User.Discriminator)
}