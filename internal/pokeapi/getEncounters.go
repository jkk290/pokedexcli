package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetEncounters(url string) (Encounters, error) {
	encountersData, exist := c.cache.Get(url)
	if exist {
		encounters := Encounters{}
		if err := json.Unmarshal(encountersData, &encounters); err != nil {
			return Encounters{}, fmt.Errorf("error converting encounters json: %v", err)
		}
		return encounters, nil
	}

	res, err := c.http.Get(url)
	if err != nil {
		return Encounters{}, fmt.Errorf("error getting encounters: %v", err)
	}
	defer res.Body.Close()

	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Encounters{}, fmt.Errorf("error converting res.Body to bytes: %v", err)
	}

	c.cache.Add(url, resBodyBytes)

	encounters := Encounters{}
	if err := json.Unmarshal(resBodyBytes, &encounters); err != nil {
		return Encounters{}, fmt.Errorf("error converting encounters json: %v", err)
	}
	return encounters, nil
}

type Encounters struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
