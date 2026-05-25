package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)
	lines := 0

	// read input line by line from stdin
	reader := bufio.NewScanner(os.Stdin)
	// for each line: split into words, count each word in a map
	for reader.Scan() {
		lines++
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

	fmt.Printf("lines: %d, unique words: %d\n", lines, len(count))
}

type WordCount struct {
	Word  string
	Count int
}

wordSort := []WordCount{}
for k, v := range count {
	wordSort = append(wordSort, WordCount{Word: k, Count: v})
	sort.Slice(wordSort, func(i, j int) bool {
		if
	})
	}
