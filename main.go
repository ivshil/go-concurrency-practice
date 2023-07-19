package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go printMessage("Hello, World!")
	fmt.Println("Goroutine started!")

	c := make(chan string)
	var wg sync.WaitGroup

	fileNames := []string{"Readme.md", "main.go", "go.mod", "main_test.go", "process_tracker.go"}
	fileNamesLen := len(fileNames)

	wg.Add(fileNamesLen)

	for i := 0; i < fileNamesLen; i++ {
		go processFile(fileNames[i], c, &wg)
	}

	go func() {
		wg.Wait() // Wait for all goroutines to complete.
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}

	fmt.Println("Well, well, well..")
	fmt.Println("End of Program")
}

func printMessage(str string) {
	fmt.Println(str)
}

func processFile(fileName string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate some processing time
	time.Sleep(12 * time.Millisecond)

	c <- fmt.Sprintf("\nProcessing is being done with file: %s\n", fileName)
}
