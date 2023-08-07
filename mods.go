package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path>")
		os.Exit(1)
	}

	path := os.Args[1]
	findFilesModifiedWithinLastWeek(path)
}

func findFilesModifiedWithinLastWeek(path string) {
	fmt.Printf("Finding files modified within the last week in: %s\n", path)

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file was modified within the last week
		modifiedTime := info.ModTime()
		oneWeekAgo := time.Now().AddDate(0, 0, -7)

		if modifiedTime.After(oneWeekAgo) {
			fmt.Println(filePath)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
