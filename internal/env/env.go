package env

import (
	"log"
	"os"
)

func GetString(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal("You need to save an api key in your .env file as API_KEY")
	}

	return val
}