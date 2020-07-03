package cities


func BuildResponseCities(file CitiesFile) []ResponseCity {
	var cities []ResponseCity
	states := BuildStatesMap(file)

	for idx := range file.Cities {
		city := file.Cities[idx]
		state := states[city.StateId]

		cities = append(cities, ResponseCity{
			Id:      city.Id,
			Name:    city.Name,
			Verbose: city.Name + ", " + state.Name,
			State:   state,
		})
	}

	return cities
}

func BuildStatesMap(file CitiesFile) map[int]State {
	states := map[int]State{}

	for idx := range file.States {
		state := file.States[idx]
		states[state.Id] = state
	}

	return states
}
