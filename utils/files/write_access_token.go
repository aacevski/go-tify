package utils

import (
	"log"

	"github.com/zalando/go-keyring"
)

func WriteAccessToken(access_token string) {
	err := keyring.Set(serviceName, accountName, access_token)
	if err != nil {
		log.Fatal(err)
	}
}
