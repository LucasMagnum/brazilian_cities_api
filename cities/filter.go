package cities

import (
	"encoding/json"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strconv"
	"strings"
	"unicode"
)

func FilterCities(search string, cities []ResponseCity) []byte {
	responseCities := []ResponseCity{}
	search = strings.ToLower(search)

	for idx := range cities {
		city := cities[idx]
		name := strconv.Quote(strings.ToLower(city.Name))

		transformer := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		normalizedName, _, _ := transform.String(transformer, name)
		normalizedSearch, _, _ := transform.String(transformer, search)

		if strings.Contains(normalizedName, normalizedSearch) {
			responseCities = append(responseCities, city)
		}
	}

	jsonResponse, _ := json.Marshal(responseCities)
	return jsonResponse
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
