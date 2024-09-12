package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("RECOVERED: %v\n", err)
		}
	}()
	cmdName, err := exec.LookPath(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var args []string
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}

	cmdArgs := append([]string{cmdName}, os.Args[2:]...)
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
