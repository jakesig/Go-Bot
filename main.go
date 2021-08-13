/* main.go
** Go Bot
** Author: Jake Sigman
** This file contains the primary code for operating the bot.
 */

package main

// Imports

import (
	"PainBot/config"
	"fmt"
)

// Main function, with prints showing status

func main() {

	// Initialize the bot

	fmt.Println("Initializing bot...")
	config.Init()

	// Initialize the server

	fmt.Println("Initializing server...")
	config.Server()

}
