package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocations(url string) (Location, error) {
	res, err := http.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("error getting locations: %v", err)
	}
	defer res.Body.Close()

	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("error converting response body to bytes: %v", err)
	}
	locations := Location{}

	if err := json.Unmarshal(resBodyBytes, &locations); err != nil {
		return Location{}, fmt.Errorf("error converting json: %v", err)
	}
	return locations, nil
}

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name *string `json:"name"`
		URL  *string `json:"url"`
	} `json:"results"`
}
