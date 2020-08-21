package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

func key() int {
	keyLength, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println("Error:", os.Args[2], "is not a valid number")
		fmt.Println(err.Error())
		return 1
	}
	if keyLength <= 0 {
		fmt.Println("Error: the key length has to be positive")
		return 1
	}

	randomBytes := make([]byte, 4096)

	// Is key file existing?
	_, err = os.Stat(os.Args[3])
	if err == nil {
		fmt.Println("Error:", os.Args[3], "already exists")
		return 1
	}

	// Create key file
	keyFile, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Println("Error: while creating the key file")
		fmt.Println(err.Error())
		return 1
	}
	defer keyFile.Close()
	keyWriter := bufio.NewWriter(keyFile)

	for i := int64(0); i < keyLength; i += 4096 {
		// Generate random bytes and write to key
		_, err = rand.Read(randomBytes)
		if err != nil {
			fmt.Println("Error: reading random bytes was not successful")
			fmt.Println(err.Error())
			return 1
		}

		bytesToWrite := int64(4096)
		if keyLength-i < 4096 {
			bytesToWrite = keyLength - i
		}
		_, err = keyWriter.Write(randomBytes[0:bytesToWrite])
		if err != nil {
			fmt.Println("Error: writing to key file was not successful")
			fmt.Println(err.Error())
			return 1
		}

		// Print a dot every MiB
		if (i+4096)%1048576 == 0 {
			fmt.Print(".")
		}
	}

	// New line after the dots
	fmt.Println()

	err = keyWriter.Flush()
	if err != nil {
		fmt.Println("Error: flushing data was not successful")
		fmt.Println(err.Error())
		return 1
	}

	return 0
}
