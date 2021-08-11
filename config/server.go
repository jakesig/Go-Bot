/* server.go
** Go Bot
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
    http.ListenAndServe(":8080", nil)
}

// Server helper function

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bot is alive!")
}