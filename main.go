package main

import (
	"encoding/json"
	"fmt"

	api "github.com/ntwaliheritier/giphy_scrapper/api"
)

func main() {
	data := api.FetchGif("bear", 3)
	b, err := json.MarshalIndent(data, "", "  ") 
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}