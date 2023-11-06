package fetchers

import (
	"io"
	"net/http"
	"strings"
)

func Play_Song(access_token string, song string) {
	context_uri := Search_Song(access_token, song)

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + access_token,
	}

	client := &http.Client{}

	url := "https://api.spotify.com/v1/me/player/play"

	req, _ := http.NewRequest("PUT", url, nil)

	req.Body = io.NopCloser(strings.NewReader(`{"context_uri": "` + context_uri + `"}`))

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if (err) != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	println(string(body))
}

func PlayUri(access_token string, track string) {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + access_token,
	}

	client := &http.Client{}

	url := "https://api.spotify.com/v1/me/player/play"

	req, _ := http.NewRequest("PUT", url, nil)

	req.Body = io.NopCloser(strings.NewReader(`{"uris": ["` + track + `"]}`))

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if (err) != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
