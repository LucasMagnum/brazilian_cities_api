package main

import (
	"encoding/json"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)
import _ "golang.org/x/text/unicode/norm"


type State struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Verbose string `json:"verbose"`
}

type City struct {
	Id int `json:"id"`
	Name string `json:"name"`
	StateId int `json:"state_id"`
}

type ResponseCity struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Verbose string `json:"verbose"`
	State State `json:"state"`
}


type CitiesFile struct {
	Cities []City `json:"cities"`
	States []State `json:"states"`
}

func main() {
	file := loadCities()
	cities := buildResponseCities(file)
	allCitiesJSON, _ := json.Marshal(cities)

	http.HandleFunc("/cities", func(writer http.ResponseWriter, request *http.Request) {
		var jsonResponse []byte

		searchParameters, _ := request.URL.Query()["q"]
		if len(searchParameters) > 0 {
			filteredCities := filterCities(searchParameters[0], cities)
			jsonResponse = filteredCities
		} else {
			jsonResponse = allCitiesJSON
		}

		writer.Header().Set("content-type", "application/json")
		writer.Write(jsonResponse)
	})

	log.Println("Server running on Port :8009")
	log.Fatal(http.ListenAndServe(":8009", nil))

}

func loadCities() CitiesFile {
	byteValue, _ := ioutil.ReadFile("./cities.json")

	file := CitiesFile{}
	json.Unmarshal(byteValue, &file)

	return file
}

func buildResponseCities(file CitiesFile) []ResponseCity {
	var cities []ResponseCity
	states := buildStatesMap(file)

	for idx := range file.Cities {
		city := file.Cities[idx]
		state := states[city.StateId]

		cities = append(cities, ResponseCity{
			Id: city.Id,
			Name: city.Name,
			Verbose: city.Name + ", " + state.Name,
			State: state,
		})
	}

	return cities
}

func buildStatesMap(file CitiesFile) map[int]State {
	states := map[int]State{}

	for idx := range file.States {
		state := file.States[idx]
		states[state.Id] = state
	}

	return states
}

func filterCities(search string, cities []ResponseCity) []byte {
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