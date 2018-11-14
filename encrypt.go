package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func checkForError(e error) {
	if e != nil && e != io.EOF {
		panic(e)
	}
}

func main() {
	totalBytesProcessed := 0
	randomByte := make([]byte, 1)
	readBytes := make([]byte, 512)

	// File needs to be specified
	if len(os.Args) > 1 {
		// Open source file
		sourceFile, err := os.Open(os.Args[1])
		checkForError(err)
		defer sourceFile.Close()

		// New buffered reader for source file
		sourceReader := bufio.NewReader(sourceFile)

		// Create destination file
		destinationFile, err := os.Create(os.Args[1] + ".fizz")
		checkForError(err)
		defer destinationFile.Close()

		fmt.Printf("Created destination file: %s\n", os.Args[1]+".fizz")

		// New buffered writer for destination file
		destinationWriter := bufio.NewWriter(destinationFile)

		// Create key file
		keyFile, err := os.Create(os.Args[1] + ".fizzkey")
		checkForError(err)
		defer keyFile.Close()

		fmt.Printf("Created key file: %s\n", os.Args[1]+".fizzkey")

		// New buffered writer for key file
		keyWriter := bufio.NewWriter(keyFile)

		// Set amount of bytes read to 1 to initially enter the loop
		amountBytesRead := 1
		for amountBytesRead > 0 {
			// Read source file
			amountBytesRead, err = sourceReader.Read(readBytes)
			checkForError(err)

			// For every read byte do the following
			for counter := 0; counter < amountBytesRead; counter++ {
				// Generate random byte
				_, err = rand.Read(randomByte)
				checkForError(err)

				// Write random byte to key file
				err = keyWriter.WriteByte(randomByte[0])
				checkForError(err)

				// Write encrypted byte to destination file
				err = destinationWriter.WriteByte(randomByte[0] ^ readBytes[counter])
				checkForError(err)

				// Print a dot every kilobyte
				totalBytesProcessed++
				if totalBytesProcessed%1024 == 0 {
					fmt.Printf(".")
				}
			}
		}

		// Flush writers
		err = keyWriter.Flush()
		checkForError(err)
		err = destinationWriter.Flush()
		checkForError(err)

		// Print finished
		fmt.Printf("\nFinished encryption\n")
	} else {
		fmt.Println("Plese specify a file")
	}

}