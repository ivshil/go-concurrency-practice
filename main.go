package main

import (
	"fmt"
	"time"
)

func main() {
	//TODO
	go printMessage("Hello, World!")
	fmt.Println("Goroutine started!")
	time.Sleep(1 * time.Second)

	fileNames := []string{"Readme.md", "main.go", "go.mod", "main_test.go", "process_tracker.go"}
	for i := 0; i < len(fileNames); i++ {
		go processFile(fileNames[i])
	}

	fmt.Println("Well, well, well..")
	time.Sleep(6 * time.Second)
	fmt.Println("End of Program")
}

func printMessage(str string) {
	fmt.Println(str)
}

func processFile(fileName string) {
	time.Sleep(3 * time.Second)
	fmt.Printf("\nProcessing is being done with file: %s\n", fileName)
	time.Sleep(2 * time.Second)
}
