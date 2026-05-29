package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetLocationList(pageUrl *string) (*LocationListResponse, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil && *pageUrl != "" {
		url = *pageUrl

	}

	cached, exist := c.cache.Get(url)
	if exist {
		var response LocationListResponse

		if err := json.Unmarshal(cached, &response); err != nil {
			return nil, err
		}

		return &response, nil
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

	marshaled, err := json.Marshal(locationList)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, marshaled)

	return &locationList, nil
}

func (c *Client) GetLocationArea(areaName string) (*LocationAreaResponse, error) {
	url := baseURL
	if areaName != "" {
		url = baseURL + "/location-area/" + areaName
	}

	cached, ok := c.cache.Get(url)
	if ok {
		var areaResponse LocationAreaResponse

		if err := json.Unmarshal(cached, &areaResponse); err != nil {
			return nil, err
		}

		return &areaResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong %w", err)
	}

	client := c.HTTPClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("something went wrong %d", resp.StatusCode)
	}

	var result LocationAreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	marshaled, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, marshaled)

	return &result, nil
}
