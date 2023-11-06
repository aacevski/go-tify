package spotify

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/aacevski/go-tify/utils/url"
	"github.com/joho/godotenv"
)

func Fetch_Access_Token(code string) string {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	godotenv.Load(".env")

	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	client_secret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	access_token := "Basic " + base64.StdEncoding.EncodeToString([]byte(client_id+":"+client_secret))

	client := &http.Client{}

	urlParams := map[string]string{
		"grant_type":   "authorization_code",
		"code":         code,
		"redirect_uri": "http://localhost:420",
	}

	url := "https://accounts.spotify.com/api/token?" + url.Build_Query(urlParams)

	req, _ := http.NewRequest("POST", url, nil)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	req.Header.Add("Authorization", access_token)

	resp, err := client.Do(req)

	if (err) != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	return result["access_token"].(string)
}
