package file

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"github.com/joaopires/hatch/internal/utils"
)

type ProcessResult struct {
	Result string
	Err    error
}

/* This approach is based on the idea of having an hash representation of each json object in the array.
   Since the key order inside a json objects matters for the hashing process, each object is decoded to a map,
   has its keys sorted and using the fmt.Sprintf method which guarantees map string representation with ordered keys,
   the map is stringified and passed to the md5 hash function. The hash is stored as a key in a map for later comparison.
   Furthermore, this method should receive a io.Reader and not be bound to a file implementation for flexibility and
   for allowing other json data sources including byte streams for better testability but to keep things more simple for
   the exercise and because of my time constraint, i'll leave it like this
*/
func Process(filePath string) ProcessResult {
	file, err := os.Open(filePath)

	if err != nil {
		return ProcessResult{
			Err: err,
		}
	}

	defer file.Close()

	jsonDecoder := json.NewDecoder(file)

	/* Should verify if the starting token is '['
	   but for the challenge scope, i'll assume that
	   the given file is always an array of json objects
	*/
	_, err = jsonDecoder.Token()

	if err != nil {
		return ProcessResult{
			Err: err,
		}
	}

	// The map containing the hashified json objects
	md5Map := make(map[string]bool)

	// while the array contains values
	for jsonDecoder.More() {
		var m map[string]interface{}

		err := jsonDecoder.Decode(&m)
		if err != nil {
			return ProcessResult{
				Err: err,
			}
		}

		// Since this is a program to compare the integrity between two json files,
		// we can use a little hack with fmt.Sprint because the string representation
		// of map is always generated with all keys ordered
		hashedObj := utils.GetMD5Hash(fmt.Sprint(m))

		/* The value is a boolean since the map entries only matter for the key
		   and the boolean type has the minimal memory allocation
		*/
		md5Map[hashedObj] = true
	}

	_, err = jsonDecoder.Token()

	if err != nil {
		return ProcessResult{
			Err: err,
		}
	}

	// Calling again the fmt.Sprint to sort all hash keys
	return ProcessResult{
		Result: fmt.Sprint(md5Map),
	}
}

func EqualSize(filepathA, filepathB string) (bool, error) {
	var fileSizeA, fileSizeB int64
	var err error

	if fileSizeA, err = getFileSize(filepathA); err != nil {
		return false, err
	}

	if fileSizeB, err = getFileSize(filepathB); err != nil {
		return false, err
	}

	return fileSizeA == fileSizeB, nil
}

func getFileSize(filepath string) (int64, error) {
	var fStats fs.FileInfo
	var err error

	if fStats, err = os.Stat(filepath); err != nil {
		return -1, err
	}

	return fStats.Size(), nil
}
