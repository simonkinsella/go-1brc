package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	FileMunchSize = 85
	BatchChanSize = 10
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
	defer file.Close()

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
				wg.Add(1)
				batchChan <- batch

				remainsLen = bytesRead - splitPoint
				remains = make([]byte, remainsLen)
				copy(remains, munch[splitPoint:])
				break
			}
		}
	}

	wg.Wait()
	close(batchChan)

	return
}

func ProcessBatch(batch []byte) {
	fmt.Print(string(batch))
}

