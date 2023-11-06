package fetchers

import (
	"encoding/json"
	"io"
	"net/http"
)

type SpotifyTrackReponse struct {
	Tracks struct {
		Href  string `json:"href"`
		Items []struct {
			Album struct {
				AlbumType string `json:"album_type"`
				Artists   []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
				AvailableMarkets string `json:"available_markets"`
				ExternalUrls     struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				ID     string `json:"id"`
				Images []struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				TotalTracks          int    `json:"total_tracks"`
				Type                 string `json:"type"`
				URI                  string `json:"uri"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			DiscNumber  int  `json:"disc_number"`
			DurationMs  int  `json:"duration_ms"`
			Explicit    bool `json:"explicit"`
			ExternalIDs struct {
				ISRC string `json:"isrc"`
			} `json:"external_ids"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string `json:"href"`
			ID          string `json:"id"`
			IsLocal     bool   `json:"is_local"`
			Name        string `json:"name"`
			Popularity  int    `json:"popularity"`
			PreviewURL  string `json:"preview_url"`
			TrackNumber int    `json:"track_number"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
		} `json:"items"`
	} `json:"tracks"`
}

func Search_Song_Raw(access_token string, song string) SpotifyTrackReponse {
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

	var trackReponse SpotifyTrackReponse

	// Read the response body into a byte slice.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle the error here.
	}

	err = json.Unmarshal(body, &trackReponse)
	if err != nil {
		// Handle the error here.
	}

	return trackReponse
}

func Search_Song(access_token string, song string) string {
	result := Search_Song_Raw(access_token, song)
	uri := ""

	for _, item := range result.Tracks.Items {
		uri = item.Album.URI
		break
	}

	return uri
}
