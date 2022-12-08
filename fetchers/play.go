package fetchers

import (
	"io/ioutil"
	"net/http"
)

func Play(access_token string) {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + access_token,
	}

	client := &http.Client{}

	url := "https://api.spotify.com/v1/me/player/play"

	req, _ := http.NewRequest("PUT", url, nil)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if (err) != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	println(string(body))
}
