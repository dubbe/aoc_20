package helpers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// check
func Check(e error) {
	if e != nil {
			panic(e)
	}
}

// readlines
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
			return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// readlines
func ReadInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
			return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			number, err := strconv.Atoi(scanner.Text())
			if err == nil {
				lines = append(lines, number)
			}
	}
	return lines, scanner.Err()
}

// readgroups
func ReadGroups(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Return nothing if at end of file and no data passed
		 if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Find the index of the input of two newlines
		if i := strings.Index(string(data), "\n\n"); i >= 0 {
			return i + 1, data[0:i], nil
		}

		// If at end of file with data return the data
		if atEOF {
			
			return len(data), data, nil
		}

		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		lines = append(lines, strings.TrimRight(strings.TrimLeft(scanner.Text(), "\n"), "\n"))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}