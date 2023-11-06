package main

import (
	"fmt"
	"log"

	"github.com/aacevski/go-tify/utils/cli"
	"github.com/aacevski/go-tify/utils/spotify"
	"github.com/zalando/go-keyring"
)

const serviceName = "gotify"
const accountName = "SpotifyAccessToken"

func main() {
	access_token, err := keyring.Get(serviceName, accountName)
	if err != nil {
		// If the error is because the item is not found, fetch a new access token.
		// Otherwise, log the error and exit.
		if err == keyring.ErrNotFound {
			fmt.Println("🔑 Generating access token...")
			code := spotify.Get_Code()
			access_token = spotify.Fetch_Access_Token(code)
			if err := keyring.Set(serviceName, accountName, access_token); err != nil {
				log.Fatalf("Failed to save access token to keyring: %v", err)
			}
		} else {
			log.Fatalf("Failed to get access token from keyring: %v", err)
		}
	} else {
		fmt.Println("🔑 Access token found in keyring!")
	}

	cli.Build()
}
