package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser
	file, err := OpenCSV("data.csv")
	if err != nil {
		//runs opencsv and handles error. theis impl always returns an error
		fmt.Printf("Error: %s", err)
		return
	}
	//use deferred function to ensure a file gets closed
	defer file.Close()

	//do something with file. normally this is where you'd do with file
}

func OpenCSV(filename string) (file *os.File, err error) {
	//opens csv and preprocesses file, note named return values
	defer func() {
		if r := recover(); r != nil {
			//main deferred error handling
			file.Close()
			err = r.(error)
		}
	}()
	file, err = os.Open(filename)
	if err != nil {
		//open file and handles errors
		fmt.Printf("failed to open file\n")
		return file, err
	}
	RemoveEmptyLines(file)
	//runs intentially broken removeEmptyLines function
	return file, err
}
func RemoveEmptyLines(f *os.File) {
	//instead of stripping empty lines, always fails here
	panic(errors.New("failed parse"))
}
