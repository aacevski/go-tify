package main

import (
	"os"

	"github.com/aacevski/go-tify/utils/cli"
	utils "github.com/aacevski/go-tify/utils/files"
	"github.com/aacevski/go-tify/utils/spotify"
)

func main() {
	access_token := ""

	if _, err := os.Stat("access_token.txt"); err == nil {
		println("🔑 Access token found!")
		access_token = utils.Read_Access_Token()
	} else {
		println("🔑 Generating access token...")
		code := spotify.Get_Code()
		access_token = spotify.Fetch_Access_Token(code)
		utils.Write_Access_Token(access_token)
	}

	cli.Build()
}
