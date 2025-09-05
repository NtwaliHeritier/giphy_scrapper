package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func FetchGif(client *http.Client, baseURL, apiKey, query string, limit int) (Response, error) {
	url := fmt.Sprintf("%s/v1/gifs/search?api_key=%s&q=%s&limit=%d",
		baseURL, apiKey, query, limit,
	)
	resp, err := client.Get(url)

	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var data Response
	
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return Response{}, err

	}
	return data, nil
}