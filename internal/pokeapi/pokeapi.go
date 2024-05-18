package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
)

type MapResult struct {
	Count     int     `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetMapFromPokeApi(url *string) (MapResult, error) {
	baseUrl := "https://pokeapi.co/api/v2/location/" // Default locations url.
	results := MapResult{}

	if url != nil {
		baseUrl = *url
	}

	// Check if we can retreive from cache
	bytes, result := c.cache.Get(baseUrl)
	if result {
		json.Unmarshal(bytes, &results)
		return results, nil
	}
	res, err := c.client.Get(baseUrl)

	if err != nil {
		return results, errors.New("failed to fetch from pokeApi")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode >= 300 {
		return results, errors.New("endpoint returned failed status code")
	}
	if err != nil {
		return results, errors.New("failed to parse body after fetching from pokeApi")
	}
	// Cache this result for future use
	c.cache.Add(baseUrl, body)

	err = json.Unmarshal(body, &results)
	if err != nil {
		return MapResult{}, err
	}

	return results, nil
}
