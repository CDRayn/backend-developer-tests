package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	exp := regexp.MustCompile(`(?mi).*error.*`)

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadSlice('\n')
		if err == nil {
			if exp.Match(line) {
				fmt.Println(string(line))
			}
		} else if err == io.EOF {
			if exp.Match(line) {
				fmt.Println(string(line))
			}
			break
		}
	}
}
