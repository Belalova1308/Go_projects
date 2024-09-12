package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func fileToMap(filepath string) (map[string]bool, error) {
	file, err := os.Open(filepath)
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make(map[string]bool)
	for scanner.Scan() {
		lines[scanner.Text()] = true
	}
	if err := scanner.Err(); err != nil {
		os.Exit(3)
	}
	return lines, err
}
func copareFiles(filepath string, linesMap map[string]bool, added bool) error {
	file, err := os.Open(filepath)
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, found := linesMap[line]
		if found != true {
			if added {
				fmt.Println("ADDED", line)
			} else {
				fmt.Println("REMOVED", line)
			}
		}
		delete(linesMap, line)
	}
	return nil
}
func main() {
	flags := flag.NewFlagSet("flags", flag.ExitOnError)
	oldFile := flags.String("old", "", "path to old")
	newFile := flags.String("new", "", "path to new")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		os.Exit(2)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("RECOVERED: %v\n", err)
		}
	}()

	oldLines, err := fileToMap(*oldFile)
	newLines, err := fileToMap(*newFile)
	if err != nil {
		os.Exit(2)
	}
	err = copareFiles(*newFile, oldLines, true)
	if err != nil {
		os.Exit(2)
	}
	err = copareFiles(*oldFile, newLines, false)
	if err != nil {
		os.Exit(2)
	}
}
