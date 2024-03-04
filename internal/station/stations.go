package station

import "strconv"

type Stations map[string]*Station

func (s Stations) Add(name, temp string) {
	t, _ := strconv.ParseFloat(temp, 64)
	station, exists := s[name]
	if !exists {
		station = &Station{}
		s[name] = station
	}
	station.AddTemp(t)
}

func (s Stations) Merge(moreStations Stations) {
	for name, station := range moreStations {
		existingStation, exists := s[name]
		if !exists {
			s[name] = station
			continue
		}
		existingStation.Min = min(station.Min, existingStation.Min)
		existingStation.Max = max(station.Max, existingStation.Max)
		existingStation.count += station.count
		existingStation.total += station.total
	}
}
