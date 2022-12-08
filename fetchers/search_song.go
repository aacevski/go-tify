package fetchers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Search_Song(access_token string, song string) string {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + access_token,
	}

	client := &http.Client{}

	url := "https://api.spotify.com/v1/search?q=" + song + "&type=track&limit=50&offset=0"

	req, _ := http.NewRequest("GET", url, nil)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if (err) != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	uri := ""

	for _, item := range result["tracks"].(map[string]interface{})["items"].([]interface{}) {
		uri = item.(map[string]interface{})["album"].(map[string]interface{})["uri"].(string)
		break
	}

	return uri
}
