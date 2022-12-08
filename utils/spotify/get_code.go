package spotify

import (
	"strings"

	"github.com/aacevski/go-tify/utils/url"
	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
)

func Get_Code() string {
	// Seperate scopes with space (encoded which is %20)
	scopes := strings.Join(Get_Scopes(), "%20")

	url_params := map[string]string{
		"client_id":     "27759326a5b8493a87ee6bcae5aae99a",
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
