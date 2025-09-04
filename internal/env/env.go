package env

import (
	"errors"
	"os"
)

func GetString(key string) (string, error) {
	val, ok := os.LookupEnv(key)

	if !ok {
		return "", errors.New("you need to save an api key in your .env file as API_KEY")
	}

	return val, nil
}