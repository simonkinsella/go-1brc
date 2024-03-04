package station

import (
	"testing"
)

type TestCase struct {
	scenario      func(s *Station)
	expectedCount int
	expectedMin,
	expectedMax,
	expectedTotal,
	expectedMean float64
}

func NewTestCase(scenario func(s *Station), expectedCount int, expectedMin, expectedMax, exptectedTotal, expectedMean float64) TestCase {
	return TestCase{
		scenario:      scenario,
		expectedCount: expectedCount,
		expectedMin:   expectedMin,
		expectedMax:   expectedMax,
		expectedTotal: exptectedTotal,
		expectedMean:  expectedMean,
	}
}

var testCases = []TestCase{
	NewTestCase(func(s *Station) {
	},
		0,
		0,
		0,
		0,
		0),
	NewTestCase(func(s *Station) {
		s.AddTemp(5)
	},
		1,
		5,
		5,
		5,
		5),
	NewTestCase(func(s *Station) {
		s.AddTemp(-5)
	},
		1,
		-5,
		-5,
		-5,
		-5),
	NewTestCase(func(s *Station) {
		s.AddTemp(-3)
		s.AddTemp(7)
		s.AddTemp(8)
	},
		3,
		-3,
		8,
		12,
		4),
}

func TestStation(t *testing.T) {

	for i, testCase := range testCases {
		s := &Station{}
		testCase.scenario(s)

		if testCase.expectedMin != s.Min {
			t.Logf("Scenario %d: Expected Station.Min to be %f (got %f)", i, testCase.expectedMin, s.Min)
			t.Fail()
		}
		if testCase.expectedMax != s.Max {
			t.Logf("Scenario %d: Expected Station.Max to be %f (got %f)", i, testCase.expectedMax, s.Max)
			t.Fail()
		}
		if testCase.expectedTotal != s.total {
			t.Logf("Scenario %d: Expected Station.total to be %f (got %f)", i, testCase.expectedTotal, s.total)
			t.Fail()
		}
		if testCase.expectedCount != s.count {
			t.Logf("Scenario %d: Expected Station.count to be %d (got %d)", i, testCase.expectedCount, s.count)
			t.Fail()
		}
		mean := s.Mean()
		if testCase.expectedMean != mean {
			t.Logf("Scenario %d: Expected to Station.Mean() to be %f (got %f)", i, testCase.expectedMean, mean)
			t.Fail()
		}
	}
}

func TestEdgeCase(t *testing.T) {
	s := &Station{
		Min:   7.3,
		Max:   36.0,
		count: 8,
		total: 174.800000,
	}
	res := s.Mean()
	expected := 21.85
	if res != expected {
		t.Logf("Expected Mean() to return %f, but got %f", expected, res)
		t.Fail()
	}
}
