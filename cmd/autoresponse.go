package cmd

import (
  "os"
  "strings"
  "github.com/bwmarrin/discordgo"
)

func Autoresponse(s *discordgo.Session, m *discordgo.MessageCreate, autoresponses map[string]string) {

  args := strings.Split(m.Content, " ")
  write := args[1]+"/"+args[2]

  f, _ := os.OpenFile("init.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

  autoresponses[args[1]] = args[2]

  if len(args) > 3 {
    for i := 3; i < len(args); i++ {
      write += " " + args[i];
    }
    key := strings.Split(write, "/")
    autoresponses[args[1]] = key[1]
  }

  f.Write([]byte("\n" + write))

  f.Close()
}