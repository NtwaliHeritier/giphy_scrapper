package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	env "github.com/ntwaliheritier/giphy_scrapper/internal/env"
)

type Response struct {
    Data []GIF `json:"data"`
}

type GIF struct {
    ID     string `json:"id"`
    URL    string `json:"url"`
	Username string `json:"username"`
	Title string `json:"title"`
}

func FetchGif(query string, limit int) Response {
	apiKey := env.GetString("API_KEY")
	url := fmt.Sprintf("https://api.giphy.com/v1/gifs/search?api_key=%s&q=%s&limit=%d", apiKey, query, limit)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data Response
	
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}
	return data
}