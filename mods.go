package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	// Command-line flags
	path := flag.String("path", ".", "The path to search recursively")
	daysAgo := flag.Int("days", 7, "Number of days ago from now to consider modification time")
	poolSize := flag.Int("poolsize", 5, "Size of the goroutine pool")

	flag.Parse()

	// Parse the path and get absolute path
	absPath, err := filepath.Abs(*path)
	if err != nil {
		fmt.Println("Error parsing the path:", err)
		return
	}

	// Get current time and time limit (days ago)
	now := time.Now()
	timeLimit := now.AddDate(0, 0, -*daysAgo)

	// Channel for collecting found files
	filesFound := make(chan string)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a pool of goroutines
	for i := 0; i < *poolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range filesFound {
				checkModifiedTime(file, timeLimit)
			}
		}()
	}

	// Traverse the directory tree recursively
	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filesFound <- path
		}
		return nil
	})

	close(filesFound)
	wg.Wait()
}

func checkModifiedTime(file string, timeLimit time.Time) {
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println("Error reading file info:", err)
		return
	}
	modTime := info.ModTime()
	if modTime.After(timeLimit) {
		fmt.Println(file)
	}
}
