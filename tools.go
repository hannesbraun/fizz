package main

import (
	"fmt"
	"io"
	"sync"
	"time"
)

var wg sync.WaitGroup

func checkForError(e error) {
	if e != nil && e != io.EOF {
		panic(e)
	}
}

func printProcessedBytes(processedBytes *int, totalBytes int64, active *bool) {
	for *active {
		// Print status
		fmt.Printf("%20d bytes processed/%20d bytes in total", *processedBytes, totalBytes)
		for counter := 0; counter < 72; counter++ {
			fmt.Printf("\b")
		}

		// Wait a second for next update
		time.Sleep(1000 * time.Millisecond)
	}

	// Update before exiting
	fmt.Printf("%20d bytes processed/%20d bytes in total", *processedBytes, totalBytes)

	// Done: tell wait group, that this is exiting
	wg.Done()
}
