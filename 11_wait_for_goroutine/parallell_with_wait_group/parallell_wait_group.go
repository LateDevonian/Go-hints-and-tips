package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	//don't need to intialise a wait group
	var wg sync.WaitGroup

	//delcare variables here to reference i outside for loop
	var i int = -1
	var file string
	//for every file added, tell wait group your waiting for one more compress operation
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			//function calls compress and ntofies wait group that it's done
			compress(filename)
			wg.Done()
		}(file)
		//trickery with paramater passing because calling goroutine in a for loop
	}
	//outer goroutine in main waits till all compressing
	//goroutines have called wg.Done
	wg.Wait()
	fmt.Printf("Compress %d files\n", i+1)
}

func compress(filename string) error {
	in, err := os.Open(filename)
	//open source filename
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	//open a destination file and gives it .gz extension
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	//compresses data and writes it to underlying file

	_, err = io.Copy(gzout, in)
	//does the copying
	gzout.Close()

	return err
}
