package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	if len(os.Args) != 3 {
		fmt.Println("how to run it : go run main.go input.txt output.txt")
		fmt.Println("You are a failure, try again")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("File not available, check the file name:", inputFile)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var processedLines []string
	linesRead := 0
	linesRemoved := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		linesRead++

		line = strings.ReplaceAll(line, "TODO:", "ACTION:")
		line = strings.ReplaceAll(line, "CLASSIFIED:", "[REDACTED]:")

		if line == strings.ToUpper(line) && line != "" {
			words := strings.Fields(strings.ToLower(line))
			for i, w := range words {
				words[i] = strings.ToUpper(string(w[0])) + w[1:]
			}
			line = strings.Join(words, " ")
		}

		if line == "" {
			linesRemoved++
			continue
		}

		processedLines = append(processedLines, line)
	}

	if linesRead == 0 {
		fmt.Println("Input file is empty. Nothing to process.")
		return
	}

	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Cannot write to output file:", outputFile)
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	writer.WriteString("SENTINEL FIELD REPORT — PROCESSED\n")

	for i, line := range processedLines {
		writer.WriteString(fmt.Sprintf("%3d: %s\n", i+1, line))
	}

	writer.Flush()
	fmt.Println("Done")

	
}