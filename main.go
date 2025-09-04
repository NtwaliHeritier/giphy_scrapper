package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	api "github.com/ntwaliheritier/giphy_scrapper/api"
)

func main() {
	gifs := []string{"bear", "mouse", "people", "computer", "football"}
	var wg sync.WaitGroup
	dataChannel := make(chan api.Response)

	startTime := time.Now()
	go func() {
		for _, gif := range gifs {
			wg.Add(1)
			go func() {
				defer wg.Done()
				data, err := api.FetchGif(gif, 3)

				if err != nil {
					log.Fatal(err)
				}

				dataChannel <- data
			}()
		}
		wg.Wait()
		close(dataChannel)
	}()

	for data := range dataChannel {
		b, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		fmt.Println("============")
	}

	fmt.Printf("It took %v to fetch results", time.Since(startTime))
}