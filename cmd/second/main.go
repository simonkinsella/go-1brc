package main

import (
	"flag"
	"fmt"
	"github.com/simonkinsella/go-1brc/internal/station"
	"io"
	"log"
	"os"
	"slices"
	"sync"
)

const (
	FileMunchSize = 1024 * 1024
	BatchChanSize = 10
	retChanSize   = 1000
)

func main() {
	// Read args
	filename := flag.String("in", "measurements.txt", "Input file")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	retChan := make(chan station.Stations, retChanSize)
	retWg := sync.WaitGroup{}

	var stations station.Stations
	go func() {
		for stationsBatch := range retChan {
			if stations == nil {
				stations = stationsBatch
			} else {
				stations.Merge(stationsBatch)
			}
			retWg.Done()
		}
	}()

	batchChan := make(chan []byte, BatchChanSize)
	go func() {
		for batch := range batchChan {
			go ProcessBatch(batch, retChan)
			retWg.Add(1)
		}
	}()

	munch := make([]byte, FileMunchSize, FileMunchSize)
	remains := make([]byte, 0, FileMunchSize)
	remainsLen := 0
	lastChunk := false
	var totalBytesRead int64 = 0
	for !lastChunk {
		bytesRead, err := file.ReadAt(munch, totalBytesRead)
		totalBytesRead += int64(bytesRead)
		if err != nil {
			if err != io.EOF {
				log.Fatalf(err.Error())
			}
			lastChunk = true
		}

		for i := bytesRead - 1; i > 0; i-- {
			if munch[i] == '\n' {
				splitPoint := i + 1
				batchLen := remainsLen + splitPoint
				batch := make([]byte, batchLen)

				copy(batch, remains)
				copy(batch[remainsLen:], munch[:splitPoint])
				batchChan <- batch

				remainsLen = bytesRead - splitPoint
				remains = make([]byte, remainsLen)
				copy(remains, munch[splitPoint:])
				break
			}
		}
	}
	close(batchChan)

	retWg.Wait()
	close(retChan)
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
		fmt.Printf("%s=%.1f/%.1f/%.1f", name, s.Min, s.Mean(), s.Max)
		if i < numStations-1 {
			fmt.Print(",\n")
		}
	}
	fmt.Println("} ")
}

func ProcessBatch(batch []byte, retCh chan station.Stations) {
	stations := station.Stations{}
	start := 0
	var name, temp string
	batchLen := len(batch)
	for i, v := range batch {

		if v == ';' {
			name = string(batch[start:i])
			start = i + 1
			continue
		}
		if v == '\n' || i+1 == batchLen {
			temp = string(batch[start:i])
			stations.Add(name, temp)
			start = i + 1
		}
	}
	retCh <- stations
}
