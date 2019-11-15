package main

//thisis the simple tool for a start that uses one goroutine

import (
	"compress/gzip"
	"io"
	"os"
)

//one goroutine needs to start others and wiat for them to finish
//so we run individual tasks inside goroutines
//so start workers, wait till done, then continue
//this function takes all files in a library and compresses them
//folder set up called exampledata with three text files in it
//run this by go run wait_for_goroutine.go exampledata/*
//it creates a compressed fversion of the files

func main() {
	//creates a list of files passed in on command line
	for _, file := range os.Args[1:] {
		compress(file)
	}
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
