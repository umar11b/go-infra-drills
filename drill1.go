package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)
	lines := 0

	for scanner.Scan() {
		lines++
		for _, word := range strings.Fields(scanner.Text()) {
			counts[strings.ToLower(word)]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}

	fmt.Printf("lines: %d, unqie words: %d\n", lines, len(counts))
}
