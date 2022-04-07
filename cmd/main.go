package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

	fileSizeA := file.GetFileSize(filePathA)


	fileSizeB := file.GetFileSize(filePathB)
	
	if  != file.GetFileSize(filePathB) {
		log.Fatal("files are not equal")
	}

	// Read files in chunks and build hashmaps

	// fileAStats, _ := fileA.Stat()
	// fileBStats, _ := fileB.Stat()

	// fileASize := fileAStats.Size()
	// fileBSize := fileBStats.Size()

	if SolutionA(filePathA, filePathB) {
		log.Print("Files are equal!!!")
	} else {
		log.Print("Files are not equal!!!")
	}
}

func SolutionA(filePathA, filePathB string) bool {
	var o1 []map[string]interface{}
	var o2 []map[string]interface{}

	fileA := readFile(filePathA)
	fileB := readFile(filePathB)

	json.Unmarshal(fileA, &o1)
	json.Unmarshal(fileB, &o2)

	hashmapA := make(map[string]bool, len(o1))
	hashmapB := make(map[string]bool, len(o1))

	for i, _ := range o1 {
		sortedObjA := GetMD5Hash(fmt.Sprint(o1[i]))
		sortedObjB := GetMD5Hash(fmt.Sprint(o2[i]))

		hashmapA[sortedObjA] = true
		hashmapB[sortedObjB] = true
	}

	return fmt.Sprint(hashmapA) == fmt.Sprint(hashmapB)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func readFile(filePath string) []byte {
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	return fileContent
}
