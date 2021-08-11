package main

import (
  "fmt"
  "GoBot/config"
)

func main() {

  fmt.Println("Initializing bot...")

  config.Init()

  fmt.Println("Initializing server...")

  config.Server()

}