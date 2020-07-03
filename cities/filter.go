package cities

import (
	"encoding/json"
	"strings"
)

func FilterCities(search string, citiesContainer []CityOutputValue) []byte {
	outputCities := []CityOutputValue{}
	search = strings.ToLower(search)
	normalizedSearch := NormalizeString(search)

	for idx := range citiesContainer {
		city := citiesContainer[idx]

		if strings.Contains(city.NormalizedName, normalizedSearch) {
			outputCities = append(outputCities, city)
		}
	}

	jsonResponse, _ := json.Marshal(outputCities)
	return jsonResponse
}
