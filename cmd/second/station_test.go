package main

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

		if testCase.expectedMin != s.min {
			t.Logf("Scenario %d: Expected Station.min to be %f (got %f)", i, testCase.expectedMin, s.min)
			t.Fail()
		}
		if testCase.expectedMax != s.max {
			t.Logf("Scenario %d: Expected Station.max to be %f (got %f)", i, testCase.expectedMax, s.max)
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
			t.Logf("Scenario %d: Expected to Station.mean() to be %f (got %f)", i, testCase.expectedMean, mean)
			t.Fail()
		}
	}
}
