package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	api "github.com/ntwaliheritier/giphy_scrapper/api"
	env "github.com/ntwaliheritier/giphy_scrapper/internal/env"
)

func main() {
	apiKey, err := env.GetString("API_KEY")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	baseURL := "https://api.giphy.com"

	fmt.Println("Enter gif suggestions")
	reader := bufio.NewReader(os.Stdin)
	inputs, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	gifs := strings.FieldsFunc(strings.TrimSpace(inputs), func(r rune) bool {
		return r == ' ' || r == ','
	})

	var wg sync.WaitGroup
	dataChannel := make(chan api.Response)

	startTime := time.Now()
	go func() {
		for _, gif := range gifs {
			wg.Add(1)
			go func(g string) {
				defer wg.Done()
				data, err := api.FetchGif(client, baseURL, apiKey, g, 3)

				if err != nil {
					log.Printf("failed to fetch %s: %v", g, err)
					return
				}

				dataChannel <- data
			}(gif)
		}
		wg.Wait()
		close(dataChannel)
	}()

	for data := range dataChannel {
		b, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			log.Printf("failed to marshal response: %v", err)
			continue
		}
		fmt.Println(string(b))
	}

	fmt.Printf("It took %v to fetch results", time.Since(startTime))
}