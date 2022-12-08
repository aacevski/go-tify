package cmd

import (
	"os"
	"strings"

	"github.com/aacevski/go-tify/fetchers"
	utils "github.com/aacevski/go-tify/utils/files"
	"github.com/aacevski/go-tify/utils/spotify"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authorizes the user to use the Spotify API",
	Long: `This command will authorize the user to use the Spotify API. It will
	create a file called access_token.txt in the root directory of the project.
	This file will contain the access token that will be used to make requests
	to the Spotify API.`,
	Run: func(cmd *cobra.Command, args []string) {

		access_token := ""

		if _, err := os.Stat("access_token.txt"); err == nil {
			// import function from utils module
			access_token = utils.Read_Access_Token()
		} else {
			code := spotify.Get_Code()
			access_token = spotify.Fetch_Access_Token(code)
			utils.Write_Access_Token(access_token)
		}

		println("✅ Successfully logged in!")

		for {
			prompt := promptui.Prompt{
				Label: "Enter a command... 🤔",
			}

			text, err := prompt.Run()

			// If you click ctrl+c, it will exit the program
			if err != nil {
				break
			}

			if strings.Contains(text, "play") {
				text = strings.Replace(text, " ", "", 1)
				text = strings.Replace(text, "play", "", 1)

				if text != "" {
					fetchers.Play_Song(access_token, text)
				} else {
					fetchers.Play(access_token)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
