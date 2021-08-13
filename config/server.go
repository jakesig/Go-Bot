/* server.go
** Pain Bot
** Author: Jake Sigman
** This file contains the code for running the bot on a server.
 */

package config

// Imports

import (
	"fmt"
	"net/http"
)

// Primary server function

func Server() {
	fmt.Println("Server is running!")
	http.HandleFunc("/", HelloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error with server!\n" + err.Error())
		return
	}
}

// Server helper function

func HelloServer(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "Bot is alive!"); err != nil {
		fmt.Println("Error with server status!\n" + err.Error())
		return
	}
}
