package spotify

import (
	"strings"

	"os"

	"github.com/aacevski/go-tify/utils/url"
	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
)

func Get_Code() string {
	scopes := strings.Join(Get_Scopes(), "%20")

	godotenv.Load(".env")
	client_id := os.Getenv("SPOTIFY_CLIENT_ID")

	url_params := map[string]string{
		"client_id":     client_id,
		"response_type": "code",
		"redirect_uri":  "http://localhost:420",
		"scope":         scopes,
	}

	url := "https://accounts.spotify.com/authorize?" + url.Build_Query(url_params)

	browser.OpenURL(url)

	prompt := promptui.Prompt{
		Label: "Enter the URL you were redirected to",
	}

	callbackUrl, _ := prompt.Run()

	return strings.Split(callbackUrl, "code=")[1]
}
