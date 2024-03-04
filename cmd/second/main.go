package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"sync"
)

const (
	FileMunchSize = 1024
	BatchChanSize = 10
)

//var Stations = map[string]*Station{}

func main() {
	// Read args
	filename := flag.String("in", "measurements.txt", "Input file")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	retCh := make(chan Stations)
	retWg := sync.WaitGroup{}

	var stations Stations
	go func() {
		for stationsBatch := range retCh {
			if stations == nil {
				stations = stationsBatch
			} else {
				stations.Merge(stationsBatch)
			}
			retWg.Done()
		}
	}()

	batchChan := make(chan []byte, BatchChanSize)
	wg := sync.WaitGroup{}

	go func() {
		for batch := range batchChan {
			go ProcessBatch(batch)
			wg.Done()
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
				//wg.Add(1)

				s := ProcessBatch(batch)

				retWg.Add(1)
				retCh <- s
				//batchChan <- batch

				remainsLen = bytesRead - splitPoint
				remains = make([]byte, remainsLen)
				copy(remains, munch[splitPoint:])
				break
			}
		}
	}

	wg.Wait()
	retWg.Wait()
	close(batchChan)
	close(retCh)
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
			fmt.Print(",\n")
		}
	}
	fmt.Println("} ")
}

func ProcessBatch(batch []byte) Stations {
	stations := Stations{}
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
	return stations
}
