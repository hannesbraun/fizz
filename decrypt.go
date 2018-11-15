package main

import (
	"bufio"
	"fmt"
	"os"
)

func decrypt() {
	totalBytesProcessed := 0
	randomByte := make([]byte, 1)
	readBytes := make([]byte, 512)

	// Getting the size
	sourceFileInfo, err := os.Stat(os.Args[2])
	checkForError(err)
	totalBytes := sourceFileInfo.Size()

	// Open source file
	sourceFile, err := os.Open(os.Args[2])
	checkForError(err)
	defer sourceFile.Close()

	// New buffered reader for source file
	sourceReader := bufio.NewReader(sourceFile)

	// Create destination file
	destinationFile, err := os.Create(os.Args[2] + ".fizz")
	checkForError(err)
	defer destinationFile.Close()

	fmt.Printf("Created destination file: %s\n", os.Args[2]+".fizz")

	// New buffered writer for destination file
	destinationWriter := bufio.NewWriter(destinationFile)

	// Open key file
	keyFile, err := os.Open(os.Args[3])
	checkForError(err)
	defer keyFile.Close()

	// New buffered reader for key file
	keyReader := bufio.NewReader(keyFile)

	// Activate status printing
	printStatus := true
	wg.Add(1)
	go printProcessedBytes(&totalBytesProcessed, totalBytes, &printStatus)

	// Set amount of bytes read to 1 to initially enter the loop
	amountBytesRead := 1
	for amountBytesRead > 0 {
		// Read source file
		amountBytesRead, err = sourceReader.Read(readBytes)
		checkForError(err)

		// For every read byte do the following
		for counter := 0; counter < amountBytesRead; counter++ {
			// Read random byte from key file
			_, err = keyReader.Read(randomByte)
			checkForError(err)

			// Write decrypted byte to destination file
			err = destinationWriter.WriteByte(randomByte[0] ^ readBytes[counter])
			checkForError(err)

			// Count the bytes processed
			totalBytesProcessed++
		}
	}

	// Stop updating on console
	printStatus = false
	wg.Wait()

	// Flush writer
	err = destinationWriter.Flush()
	checkForError(err)

	// Print finished
	fmt.Printf("\nFinished decryption\n")

}
