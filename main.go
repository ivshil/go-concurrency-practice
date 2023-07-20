package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printMessage("Hello, World!", &wg)
	fmt.Println("Goroutine started!")

	c := make(chan string)
	fileNames := []string{"Readme.md", "main.go", "go.mod", "main_test.go", "process_tracker.go"}
	fileNamesLen := len(fileNames)

	for i := 0; i < fileNamesLen; i++ {
		wg.Add(2)
		go processFile(fileNames[i], c, &wg)
		go printChannel(c, &wg)
	}

	wg.Wait()
	close(c)
	fmt.Println("End of Program")
}

func printChannel(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-c)
}

func printMessage(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
}

func processFile(fileName string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(12 * time.Millisecond)

	c <- fmt.Sprintf("\nProcessing is being done with file: %s\n", fileName)
}
