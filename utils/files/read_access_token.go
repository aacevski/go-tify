package utils

import (
	"bufio"
	"log"
	"os"
)

func Read_Access_Token() string {
	file, err := os.Open("access_token.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	return scanner.Text()
}
