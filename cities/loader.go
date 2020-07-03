package cities

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type stateInputValue struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Verbose string `json:"verbose"`
}

type cityInputValue struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	StateId int    `json:"state_id"`
}

type CitiesInputValues struct {
	Cities []cityInputValue  `json:"cities"`
	States []stateInputValue `json:"states"`
}

func LoadCitiesFromJSONFile() CitiesInputValues {
	byteValue, err := ioutil.ReadFile("./cities/data/cities.json")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to load file")
	}

	values := CitiesInputValues{}
	json.Unmarshal(byteValue, &values)

	return values
}
