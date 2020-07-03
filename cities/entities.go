package cities

type State struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Verbose string `json:"verbose"`
}

type City struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	StateId int    `json:"state_id"`
}

type ResponseCity struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Verbose string `json:"verbose"`
	State   State  `json:"state"`
}

type CitiesFile struct {
	Cities []City  `json:"cities"`
	States []State `json:"states"`
}
