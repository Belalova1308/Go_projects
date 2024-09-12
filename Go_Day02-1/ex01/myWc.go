package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func countLines(filepath string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filepath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	lineCount := 0
	for {
		_, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineCount++
	}
	fmt.Println(lineCount)
}
func countSymbols(filepath string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filepath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	symbolCount := 0
	for {
		str, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		symbolCount += utf8.RuneCountInString(string(str))
	}
	fmt.Println(symbolCount)
}
func countWords(filepath string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filepath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	wordsCount := 0
	for {
		str, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		words := strings.Fields(string(str))
		wordsCount += len(words)
	}
	fmt.Println(wordsCount)
}
func main() {
	var l, m, w bool
	flag.BoolVar(&l, "l", false, "only file")
	flag.BoolVar(&m, "m", false, "only directories")
	flag.BoolVar(&w, "w", true, "only symlinks")
	flag.Parse()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("RECOVERED: %v\n", err)
		}
	}()
	args := flag.Args()
	if len(args) == 0 {
		os.Exit(1)
	}
	allFiles := strings.Join(args, " ")
	filepaths := strings.Split(allFiles, "\t")
	var wg sync.WaitGroup
	if l {
		wg.Add(1)
		for i := range filepaths {
			go countLines(args[i], &wg)
		}
	} else if m {
		wg.Add(1)
		for i := range args {
			go countSymbols(args[i], &wg)
		}
	} else if w {
		wg.Add(1)
		for i := range args {
			go countWords(args[i], &wg)
		}
	}
	wg.Wait()
}
