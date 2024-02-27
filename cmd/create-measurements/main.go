package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"time"
)

const (
	minRecords              = 1
	maxRecords              = 1000000000
	progressInterval        = 50000000
	temperatureStdDeviation = 10
)

func main() {
	filename := flag.String("name", "measurements.txt", "name of file to create")
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("Usage: create_measurements <number of records to create>")
	}

	numRecords, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalf("Usage: create_measurements <number of records to create>")
	}

	if numRecords < minRecords || numRecords > maxRecords {
		log.Fatalf("Number of records should be between %d and %d", minRecords, maxRecords)
	}

	f, err := os.Create(*filename)
	if err != nil {
		log.Fatalf(err.Error())
	}
	w := bufio.NewWriter(f)

	numStations := len(weatherStations)
	startTime := time.Now()

	for i := range numRecords {
		if i > 0 && i%progressInterval == 0 {
			log.Printf("Wrote %d measurements in %.3fs", i, time.Since(startTime).Seconds())
		}
		station := weatherStations[rand.IntN(numStations)]
		w.WriteString(fmt.Sprintf("%s;%.1f\n", station.name, station.measurement()))
	}

	w.Flush()

	log.Printf("Created file with %d measurements in %.3fs", numRecords, time.Since(startTime).Seconds())
}
