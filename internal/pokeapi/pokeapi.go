package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (c *Client) ExploreMap(name string) ([]string, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s-area", name)
	locInfo := LocationInfo{}

	bytes, result := c.cache.Get(url)
	if result {
		err := json.Unmarshal(bytes, &locInfo)
		if err != nil {
			return nil, errors.New("error while retrieving from cache")
		}
		return locInfo.extractPokemon(), nil
	}
	res, err := c.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error while fetching url %s", url)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode >= 300 {
		return nil, errors.New("endpoint returned failed status code")
	}
	if err != nil {
		return nil, errors.New("failed to parse body after fetching from pokeApi")
	}
	c.cache.Add(url, body)
	err = json.Unmarshal(body, &locInfo)
	if err != nil {
		return nil, err
	}
	return locInfo.extractPokemon(), nil

}
