package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Wrong number of arguments")
		fmt.Println("Generate a key with:", os.Args[0], "key <keylen> key.fizzkey")
		fmt.Println("XOR with:", os.Args[0], "xor file key.fizzkey")
		os.Exit(1)
	}

	var retVal int

	if strings.ToLower(os.Args[1]) == "xor" {
		retVal = xor()
	} else if strings.ToLower(os.Args[1]) == "key" {
		retVal = key()
	} else {
		fmt.Println("Unknown mode:", os.Args[1])
		os.Exit(1)
	}

	if retVal != 0 {
		os.Exit(1)
	}
}
