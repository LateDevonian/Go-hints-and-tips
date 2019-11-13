package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	//use a waitgroup to monitor a group of goroutines
	var wg sync.WaitGroup

	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	fmt.Println("Words that appear more than once:")
	w.Lock()
	for word, count := range w.found {
		if count > 1 {
			//at the end of the program, print what i found
			fmt.Println("%s: %d\n", word, count)
		}
	}
	w.Unlock()
	//lock and unlock added to stop race condition
}

//track words in a struct. using a struct makes any refactor easier
//could also use a type
type words struct {
	sync.Mutex
	//word struct inherits mutex lock
	found map[string]int
}

//creates new word instance
func newWords() *words {
	return &words{found: map[string]int{}}
}

//track number of times you've seen the word
func (w *words) add(word string, n int) {
	w.Lock()
	//lock the object, modify map, unlock object
	defer w.Unlock()
	count, ok := w.found[word]
	//if word isnt' already tracked, add it, otherwise
	//increment the count
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	//opens a file, parse contest, count words that appear
	//copy function does all the copying for one
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	//use scanner for parsing file
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
