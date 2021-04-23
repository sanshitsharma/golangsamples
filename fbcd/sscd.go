package main

import (
	"fmt"
	"github.com/sanshitsharma/golangsamples/fbcd/matcher"
	"github.com/sanshitsharma/golangsamples/fbcd/matcher/naive"
	"github.com/sanshitsharma/golangsamples/fbcd/util"
	"io"
	"log"
	"os"
	"strings"
)

func readFromStream(buf *util.Queue, stream *strings.Reader, numBytes uint32) (uint16, error) {
	var bytesRead uint16
	for i := 0; i < int(numBytes); i++ {
		datum, err := stream.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return bytesRead, err
		}

		bytesRead++
		if err = buf.Enqueue(datum); err != nil {
			return bytesRead, err
		}
	}

	return bytesRead, nil
}

func main() {
	// Create the buffers
	sw := util.InitSlidingWindow()

	test := `she sells sea shells on the sea shore`
	fmt.Println("Compressing:", test)

	reader := strings.NewReader(test)
	if reader == nil {
		log.Fatalf("failed to open stream reader")
	}

	// Read LookAheadSize bytes into the LookAheadBuffer
	bytesRead, err := readFromStream(sw.LookAhead, reader, sw.LookAhead.GetSize())
	if err != nil {
		log.Fatalf("failed to read from stream. err: %v\n", err)
	}
	if bytesRead == 0 {
		log.Println("empty stream.. nothing to encode. DONE!!")
		os.Exit(0)
	}

	// fmt.Printf("Sliding Window: %v\n", sw)

	// Create our pattern searcher
	searcher := naive.Create()

	for !sw.LookAhead.IsEmpty() {
		match, err := searcher.GetLongest(sw)
		if err != nil {
			log.Fatalf("failed to find longest match. err: %v\n", err)
		}
		fmt.Println(match)

		// Depending on how many bytes were compressed, the following two actions need to be performed
		// 1. Shift LongestMatch length bytes from LookAheadBuffer to Search Buffer
		nextReadLen := match.Length
		if match.Type == matcher.UNCOMPRESSED {
			nextReadLen = 1
		}

		//fmt.Println("nextReadLen:", nextReadLen)

		err = searcher.ShiftBytes(sw, nextReadLen)
		if err != nil {
			log.Fatalf("failed to shift bytes from look ahead to search. err:%v\n", err)
		}
		// 2. Read LongestMatch length bytes from io.Reader into the LookAheadBuffer
		_, err = readFromStream(sw.LookAhead, reader, uint32(nextReadLen))
		if err != nil {
			log.Fatalf("failed to read %v bytes from stream. err: %v\n", nextReadLen, err)
		}

		//fmt.Printf("After shift and read Sliding Window: %v\n", sw)
	}

	/*
	fmt.Println("Find Longest Common Match b/w Search:", sw.Search, " and LookAhead:", sw.LookAhead)

	matchLens := []int{1, 3, 8, 15}
	for _, matchLen := range matchLens {
		readBytes := func(count int) {
			fmt.Println("*********************************************************")
			fmt.Printf("Reading %d more bytes\n", count)
			// First move from LookAhead into Search
			lookAheadReadIdx := sw.LookAhead.Head
			for i := 0; i < count; i++ {
				sw.Search.GetBuffer[sw.Search.Head] = sw.LookAhead.GetBuffer[lookAheadReadIdx]
				sw.Search.IncrementIndex(util.SearchBufferSize)
				lookAheadReadIdx = (lookAheadReadIdx+1)%util.LookAheadBufferSize
			}

			fmt.Printf("######## Modified Search GetBuffer: %v ########\n", sw.Search)

			// Now move from test into LookAhead
			for i := 0; i < count; i++ {
				fmt.Printf("writing %v at look ahead index %v\n", string(test[testIdx]), sw.LookAhead.Head)
				sw.LookAhead.GetBuffer[sw.LookAhead.Head] = test[testIdx]
				sw.LookAhead.IncrementIndex(util.LookAheadBufferSize)
				testIdx++
				if int(testIdx) >= len(test) {
					fmt.Printf("testIdx: %v\n", testIdx)
					break
				} else {
					fmt.Printf("testIdx: %v, next char: %v\n", testIdx, string(test[testIdx]))
				}
				fmt.Printf("Look ahead: %v\n", sw.LookAhead)
			}
			fmt.Println("*********************************************************")
		}
		readBytes(matchLen)
		fmt.Println("Find Longest Common Match b/w Search:", sw.Search, " and LookAhead:", sw.LookAhead)
	}
	 */
}
