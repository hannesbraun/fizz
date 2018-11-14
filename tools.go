package main

import (
   "io"
)

func checkForError(e error) {
	if e != nil && e != io.EOF {
		panic(e)
	}
}
