package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	mergeChunks("mediatest/uploads/1", "merged.mp4")
}

func extractNumber(fileName string) int {
	re := regexp.MustCompile(`\d+`)
	numStr := re.FindString(filepath.Base(fileName))
	num, err := strconv.Atoi(numStr)

	if err != nil {
		return -1
	}

	return num
}

func mergeChunks(inputDir, outputFile string) error {
	chunks, err := filepath.Glob(filepath.Join(inputDir, "*.chunk"))

	if err != nil {
		return fmt.Errorf("Failed to find chunks: %v", err)
	}
	sort.Slice(chunks, func(i, j int) bool {
		return extractNumber(chunks[i]) < extractNumber(chunks[j])
	})

	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("Failed to create output file: %v", err)
	}

	defer output.Close()

	for _, chunk := range chunks {
		input, err := os.Open(chunk)
		if err != nil {
			return fmt.Errorf("Failed to open chunk file: %v", err)
		}

		_, err = output.ReadFrom(input)
		if err != nil {
			return fmt.Errorf("Failed to write chunk %s to merged file: %v", chunk, err)
		}
		input.Close()
	}
	return nil
}
