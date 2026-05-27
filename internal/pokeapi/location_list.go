package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationListResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationList(pageUrl *string) (*LocationListResponse, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var locationList LocationListResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationList); err != nil {
		return nil, err
	}
	return &locationList, nil
}
