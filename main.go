package main

import (
	"encoding/json"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
)
import _ "golang.org/x/text/unicode/norm"
import "github.com/LucasMagnum/brazilian_cities_api/cities"


func main() {
	file := cities.LoadCities()
	cities := cities.BuildResponseCities(file)
	allCitiesJSON, _ := json.Marshal(cities)

	http.HandleFunc("/cities", func(writer http.ResponseWriter, request *http.Request) {
		var jsonResponse []byte

		searchParameters, _ := request.URL.Query()["q"]
		if len(searchParameters) > 0 {
			filteredCities := cities.FilterCities(searchParameters[0], cities)
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