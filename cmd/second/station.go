package main

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
		existingStation.min = min(station.min, existingStation.min)
		existingStation.max = max(station.max, existingStation.max)
		existingStation.count += station.count
		existingStation.total += station.total
	}
}

type Station struct {
	count                 int
	min, max, total, mean float64
}

func (s *Station) AddTemp(t float64) {
	if s.count == 0 {
		s.min = t
		s.max = t
	} else {
		if t < s.min {
			s.min = t
		}
		if t > s.max {
			s.max = t
		}
	}
	s.count++
	s.total += t
}

func (s *Station) Mean() float64 {
	if s.count == 0 {
		return 0
	}
	return s.total / float64(s.count)
}
