package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)

	// read input line by line from stdin
	reader := bufio.NewScanner(os.Stdin)
	// for each line: split into words, count each word in a map
	for reader.Scan() {
		line := reader.Text()
		words := strings.Fields(line)

		for _, word := range words {
			count[word]++
		}
	}
	// also count total lines
	// after the loop: check if reading errored
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}
	// print: number of lines, number of unique words
	line := 0
	line++

	fmt.Printf("lines: %d, unique words: %d\n", line, len(count))
}
