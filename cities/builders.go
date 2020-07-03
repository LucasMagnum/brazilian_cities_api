package cities

func BuildOutputCities(citiesInputValues CitiesInputValues) []CityOutputValue {
	var cities []CityOutputValue
	states := BuildStatesMap(citiesInputValues.States)

	for idx := range citiesInputValues.Cities {
		city := citiesInputValues.Cities[idx]
		state := states[city.StateId]

		cities = append(cities, BuildOutputCity(city, state))
	}

	return cities
}

func BuildOutputCity(city cityInputValue, state StateOutputValue) CityOutputValue {
	return CityOutputValue{
		Id:             city.Id,
		Name:           city.Name,
		NormalizedName: NormalizeString(city.Name),
		Verbose:        city.Name + ", " + state.Name,
		State:          state,
	}
}

func BuildStatesMap(states []stateInputValue) map[int]StateOutputValue {
	mapping := map[int]StateOutputValue{}

	for idx := range states {
		state := states[idx]
		mapping[state.Id] = StateOutputValue{
			Id:      state.Id,
			Name:    state.Name,
			Verbose: state.Verbose,
		}
	}

	return mapping
}
