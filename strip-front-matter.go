package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Check if first line is frontmatter delimiter
	if !scanner.Scan() {
		return
	}

	firstLine := strings.TrimSpace(scanner.Text())

	// Determine delimiter type
	var delimiter string
	if firstLine == "+++" {
		delimiter = "+++"
	} else if firstLine == "---" {
		delimiter = "---"
	} else {
		// No frontmatter delimiter, print first line and rest of content
		fmt.Println(scanner.Text())
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		return
	}

	// Skip lines until we find the closing delimiter
	foundClosing := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == delimiter {
			foundClosing = true
			break
		}
	}

	// If we found closing delimiter, print the rest
	if foundClosing {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
