package fetchers

import (
	"net/url"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

func Search_Song_Rows(access_token string, song string) []table.Row {
	result := Search_Song_Raw(access_token, url.QueryEscape(song))

	var rows []table.Row

	for index, item := range result.Tracks.Items {
		rows = append(rows, table.Row{strconv.Itoa(index), item.Name, item.Artists[0].Name, item.URI})
	}

	return rows
}
