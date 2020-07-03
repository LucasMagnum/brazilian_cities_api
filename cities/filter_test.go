package cities

import "testing"


func benchmarkFilterUsingArray(*testing.B) {
	cities := LoadCities()
	responseCities := BuildResponseCities(cities)

	searchWords := []string{"sao", "são josé", "a", "paUlo", "mara", "gov", "Governador"}

	for idx := range searchWords {
		FilterCities(searchWords[idx], responseCities)
	}
}