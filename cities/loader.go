package cities

import (
	"encoding/json"
	"io/ioutil"
)

func LoadCities() CitiesFile {
	byteValue, _ := ioutil.ReadFile("./data/cities.json")

	file := CitiesFile{}
	json.Unmarshal(byteValue, &file)

	return file
}