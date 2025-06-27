package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	inputFiles := []string{"server1.log", "server2.log", "server3.log"}
	err := ProcessLogs(inputFiles, "errors.log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Errors extracted successfully.")
}

func ProcessLogs(inputFiles []string, outputFile string) error {
	var wg sync.WaitGroup
	errorChan := make(chan string, 100)

	writerDone := make(chan struct{})
	go func() {
		err := writeErrorsToFile(outputFile, errorChan)
		if err != nil {
			log.Printf("Failed to write to output file: %v", err)
		}
		close(writerDone)
	}()

	for _, file := range inputFiles {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			if err := extractErrors(filename, errorChan); err != nil {
				log.Printf("Error processing file %s: %v", filename, err)
			}
		}(file)
	}

	wg.Wait()
	close(errorChan)

	<-writerDone
	return nil
}

func extractErrors(filename string, errorChan chan<- string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			errorChan <- line
		}
	}
	return scanner.Err()
}

func writeErrorsToFile(outputFile string, errorChan <-chan string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for errLine := range errorChan {
		_, err := writer.WriteString(errLine + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
