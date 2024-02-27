package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var stations = map[string]*station{}

func main() {
	// Read args
	filename := flag.String("in", "measurements.txt", "Input file")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf(err.Error())
	}

	//Open input file, scan lines and add stations and temps
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")
		temp, _ := strconv.ParseFloat(parts[1], 64)
		s, exists := stations[parts[0]]
		if !exists {
			stations[parts[0]] = &station{}
			s = stations[parts[0]]
		}
		s.AddTemp(temp)
	}

	// Sort station names
	names := make([]string, len(stations))
	i := 0
	for name, _ := range stations {
		names[i] = name
		i++
	}
	slices.Sort(names)

	// Output results
	numStations := len(names)

	fmt.Print("{")
	for i, name := range names {
		s := stations[name]
		fmt.Printf("%s=%.1f/%.1f/%.1f", name, s.min, s.Mean(), s.max)
		if i < numStations-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("} ")
}
