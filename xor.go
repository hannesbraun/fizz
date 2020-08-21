package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func xor() int {
	fileByte := make([]byte, 1)
	keyByte := make([]byte, 1)

	// Ensure key size is greater than zero
	keyInfo, err := os.Stat(os.Args[3])
	if err != nil {
		fmt.Println("Error: getting the file size for", os.Args[2], "was not successful")
		fmt.Println(err.Error())
		return 1
	} else if keyInfo.Size() <= 0 {
		fmt.Println("Error: the key size has to be greater than zero")
		return 1
	}

	// Open file
	file, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println("Error: opening", os.Args[2], "was not successful")
		fmt.Println(err.Error())
		return 1
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)

	// Open key
	key, err := os.Open(os.Args[3])
	if err != nil {
		fmt.Println("Error: opening", os.Args[3], "was not successful")
		fmt.Println(err.Error())
		return 1
	}
	defer key.Close()
	keyReader := bufio.NewReader(key)

	// Destination filename
	var destFilename string
	if strings.HasSuffix(os.Args[2], ".fizz") {
		destFilename = os.Args[2][0 : len(os.Args[2])-5]
	} else {
		destFilename = os.Args[2] + ".fizz"
	}

	// Is destination file existing?
	_, err = os.Stat(destFilename)
	if err == nil {
		fmt.Println("Error:", destFilename, "already exists")
		return 1
	}

	// Create destination file
	destFile, err := os.Create(destFilename)
	if err != nil {
		fmt.Println("Error: creating", destFilename, "was not successful")
		fmt.Println(err.Error())
		return 1
	}
	defer destFile.Close()
	destWriter := bufio.NewWriter(destFile)
	fmt.Println("Created destination file:", destFilename)

	i := int64(0)
	for {
		// Read source file
		_, err = fileReader.Read(fileByte)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: reading", os.Args[2], "was not successful")
			fmt.Println(err.Error())
			return 1
		}

		// Read key file
		_, err = keyReader.Read(keyByte)
		if err == io.EOF {
			_, err = key.Seek(0, io.SeekStart)
			if err != nil {
				fmt.Println("Error: seeking", os.Args[3], "was not successful")
				fmt.Println(err.Error())
				return 1
			}
			_, err = keyReader.Read(keyByte)
			if err != nil {
				fmt.Println("Error: reading", os.Args[3], "was not successful")
				fmt.Println(err.Error())
				return 1
			}
		} else if err != nil {
			fmt.Println("Error: reading", os.Args[3], "was not successful")
			fmt.Println(err.Error())
			return 1
		}

		// XOR and write
		destWriter.Write([]byte{fileByte[0] ^ keyByte[0]})
		if err != nil {
			fmt.Println("Error: writing to", destFilename, "was not successful")
			fmt.Println(err.Error())
			return 1
		}

		i++
		// Print a dot every MiB
		if i%1048576 == 0 {
			fmt.Print(".")
		}
	}

	// Print new line after dots
	fmt.Println()

	err = destWriter.Flush()
	if err != nil {
		fmt.Println("Error: flushing data was not successful")
		fmt.Println(err.Error())
		return 1
	}

	return 0
}
