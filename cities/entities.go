package cities

type StateOutputValue struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Verbose string `json:"verbose"`
}

type CityOutputValue struct {
	Id             int              `json:"id"`
	Name           string           `json:"name"`
	Verbose        string           `json:"verbose"`
	NormalizedName string           `json:"normalized_name"`
	State          StateOutputValue `json:"state"`
}
