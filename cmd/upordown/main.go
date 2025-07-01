package main

import "os"

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "track":
		// Store the website URL, and begin checking it
	case "untrack":
		// Stop tracking the website URL
	}
}
