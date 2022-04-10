package main

import (
	"log"
	"os"
	"time"

	"github.com/joaopires/hatch/internal/file"
)

func main() {
	invalidArgsErrorMessage := "two valid filepaths must be passed as arguments"

	if len(os.Args) < 3 {
		log.Fatal(invalidArgsErrorMessage)
	}

	filePathA := os.Args[1]
	filePathB := os.Args[2]

	if filePathA == "" || filePathB == "" {
		log.Fatal(invalidArgsErrorMessage)
	}

	if equal, err := file.EqualSize(filePathA, filePathB); err != nil || !equal {
		if err != nil {
			log.Fatal(err)
		}

		log.Fatal("files are not equal!!!")
	}

	fileAProcessChan := make(chan file.ProcessResult)
	fileBProcessChan := make(chan file.ProcessResult)

	defer func() {
		close(fileAProcessChan)
		close(fileBProcessChan)
	}()

	start := time.Now()

	log.Println("starting the comparing process...")

	go func() {
		fileAProcessChan <- file.Process(filePathA)

	}()

	go func() {
		fileBProcessChan <- file.Process(filePathB)
	}()

	procResA := <-fileAProcessChan
	procResB := <-fileBProcessChan

	log.Printf("the comparing process took: %f seconds\n", time.Since(start).Seconds())

	if procResA.Err != nil {
		log.Fatalf("error in file: %s, %s", filePathA, procResA.Err)
	}

	if procResB.Err != nil {
		log.Fatalf("error in file: %s, %s", filePathB, procResB.Err)
	}

	if procResA.Result == procResB.Result {
		log.Print("files are equal!!!")
	} else {
		log.Print("files are not equal!!!")
	}
}
