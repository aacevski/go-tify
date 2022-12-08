package utils

import (
	"log"
	"os"
)

func Write_Access_Token(access_token string) {
	f, err := os.Create("access_token.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString(access_token)
}
