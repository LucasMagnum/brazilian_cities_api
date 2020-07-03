package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)
import "github.com/lucasmagnum/brazilian_cities_api/cities"

func main() {
	citiesInputValues := cities.LoadCitiesFromJSONFile()
	responseCities := cities.BuildOutputCities(citiesInputValues)
	allCitiesJSON, _ := json.Marshal(responseCities)

	http.HandleFunc("/cities", func(writer http.ResponseWriter, request *http.Request) {
		var jsonResponse []byte

		searchParameters, _ := request.URL.Query()["q"]
		if len(searchParameters) > 0 {
			filteredCities := cities.FilterCities(searchParameters[0], responseCities)
			jsonResponse = filteredCities
		} else {
			jsonResponse = allCitiesJSON
		}

		writer.Header().Set("access-control-allow-origin", "*")
		writer.Header().Set("access-control-allow-methods", "GET,OPTIONS")
		writer.Header().Set("content-type", "application/json")
		writer.Write(jsonResponse)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8009"
	}

	log.Printf("Server running on Port :%s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
