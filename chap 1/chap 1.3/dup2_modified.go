package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	counts := map[string]map[string]int{}
	if len(files) == 0 {
		fmt.Println("There are no arguments")
	} else {
		// Opens each file and counts the lines and couples it with the appropriate filename it occurs in
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// Scans the map and prints all lines that appear two or more times and the all files that it appears in
	for line, filenames := range counts {
		if len(filenames) > 1 {
			arr := make([]string, 0)
			for filename := range filenames {
				arr = append(arr, filename)
			}
			fmt.Printf("%v\t%s\n", arr, line)
		}
	}
}
func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// Checking if the nested map has been initialised
		val, ok := counts[input.Text()]
		if ok {
			// Increments the nested map
			val[f.Name()]++
			counts[input.Text()] = val
		} else {
			// Initialises the nested map sets it the value of the filename
			counts[input.Text()] = map[string]int{f.Name(): 1}
		}
	}
}
