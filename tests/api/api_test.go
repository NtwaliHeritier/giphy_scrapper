package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/ntwaliheritier/giphy_scrapper/api"
)

func TestFetchGif(t *testing.T) {
	// Mock response
	mockResp := api.Response{
		Data: []api.GIF{
			{ID: "123", URL: "https://mock/gif", Username: "tester", Title: "Funny"},
		},
	}

	// Spin up fake HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/gifs/search" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer ts.Close()

	client := ts.Client()
	apiKey := "fake-key"

	resp, err := api.FetchGif(client, ts.URL, apiKey, "test", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 GIF, got %d", len(resp.Data))
	}
	if resp.Data[0].ID != "123" {
		t.Errorf("expected ID 123, got %s", resp.Data[0].ID)
	}
}