package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filesWithLine := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			accountLines(f, filesWithLine)
			f.Seek(0, 0)
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println("Line occurs in the following files:")
			fmt.Println(strings.Join(filesWithLine[line], " "))
		}
	}
}

func accountLines(f *os.File, filesWithLine map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		found := false
		for _, v := range filesWithLine[input.Text()] {
			if f.Name() == v {
				found = true
				break
			}
		}
		if !found {
			filesWithLine[input.Text()] = append(filesWithLine[input.Text()], f.Name())
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Leaves errors on read
}
