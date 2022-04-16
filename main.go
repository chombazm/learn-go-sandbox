package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filename := os.Args[1];

	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error while opening the file %s\n", filename)
		os.Exit(1)
		return
	}

	contents := bufio.NewScanner(file);

	for contents.Scan() {
		counts[contents.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d", line, n)
		}
	}
} 