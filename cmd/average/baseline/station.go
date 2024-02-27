package main

type station struct {
	count                 int
	min, max, total, mean float64
}

func (s *station) AddTemp(t float64) {
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

func (s *station) Mean() float64 {
	if s.count == 0 {
		return 0
	}
	return s.total / float64(s.count)
}
