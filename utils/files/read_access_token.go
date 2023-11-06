package utils

import (
	"log"

	"github.com/zalando/go-keyring"
)

func ReadAccessToken() string {
	access_token, err := keyring.Get(serviceName, accountName)
	if err != nil {
		log.Fatal(err)
	}
	return access_token
}