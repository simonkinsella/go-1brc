package station

type Station struct {
	count           int
	Min, Max, total float64
}

func (s *Station) AddTemp(t float64) {
	if s.count == 0 {
		s.Min = t
		s.Max = t
	} else {
		s.Min = min(s.Min, t)
		s.Max = max(s.Max, t)
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
