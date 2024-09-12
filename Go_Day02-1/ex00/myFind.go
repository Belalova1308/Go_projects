package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var f, d, sl bool
	var ext string
	flag.BoolVar(&f, "f", false, "only file")
	flag.BoolVar(&d, "d", false, "only directories")
	flag.BoolVar(&sl, "sl", false, "only symlinks")
	flag.StringVar(&ext, "ext", "", "extantion")
	flag.Parse()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("RECOVERED: %v\n", err)
		}
	}()
	if ext != "" && !f {
		os.Exit(1)
	}
	if !f && !d && !sl {
		f, sl, d = true, true, true
	}
	args := flag.Args()
	root := args[0]
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			os.Exit(1)
		}
		if d && info.IsDir() {
			fmt.Println(path)
		} else if sl && info.Mode()&os.ModeSymlink != 0 {
			linkDest, err := os.Readlink(path)
			if err != nil {
				fmt.Printf("%s -> [broken]\n", path)
			} else {
				fmt.Printf("%s -> %s\n", path, linkDest)
			}
		} else if f && info.Mode().IsRegular() {
			if ext != "" {
				if filepath.Ext(path) == "."+ext {
					fmt.Println(path)
				}
			} else {
				fmt.Println(path)
			}
		}
		return nil
	})
	if err != nil {
		os.Exit(1)
	}
}
