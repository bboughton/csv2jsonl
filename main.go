package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: input file required")
		os.Exit(1)
	}
	inputPath := os.Args[1]

	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	encoder := json.NewEncoder(os.Stdout)
	for line := range readFile(inputFile) {
		encoder.Encode(line)
	}
}

func readFile(file io.Reader) <-chan map[string]interface{} {
	lines := make(chan map[string]interface{})
	go func() {
		r := csv.NewReader(file)
		props, err := r.Read()
		if err != nil {
			close(lines)
			return
		}
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				continue // skip this row
			}

			line := make(map[string]interface{})
			for i, value := range record {
				line[props[i]] = value
			}
			lines <- line
		}
		close(lines)
	}()

	return lines
}
