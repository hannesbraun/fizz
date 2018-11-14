package main

import (
	"fmt"
	"os"
)

func main() {
	// File needs to be specified
	if len(os.Args) > 2 {
		if os.Args[1] == "encrypt" {
			encrypt()
		} else if os.Args[1] == "decrypt" {
			if len(os.Args) > 3 {
				decrypt()
			} else {
				fmt.Println("Please specify the encrypted file and the key")
			}
		}
	} else {
		fmt.Println("Plese specify the mode and the file(s)")
	}

}
