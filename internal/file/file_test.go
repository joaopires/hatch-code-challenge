package file

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"testing"

	"github.com/joaopires/hatch/internal/utils"
)

// Valid json files
const testFilepathA = "../../resources/equal/unordered_small_file_before_backup.json"
const testFilepathB = "../../resources/equal/unordered_small_file_after_backup.json"
const testFilepathC = "../../resources/equal/unordered_medium_file_before_backup.json"

// Invalid json files
const testFilepathMalformedBegin = "../../resources/invalid/malformed_begin.json"
const testFilepathMalformedEnd = "../../resources/invalid/malformed_end.json"
const testFilepathInvalidObject = "../../resources/invalid/invalid_obj.json"

func TestProcess(t *testing.T) {
	target := "Process"

	// Lame way of mocking the test result...
	jsonArray := []map[string]interface{}{
		{
			"id":   "jhasdad",
			"name": "test json",
			"obj": map[string]interface{}{
				"array": []int{1, 5, 6, 7, 8},
				"age":   19,
			},
		},
		{
			"id":   "wqweq",
			"name": "test json 2",
		},
	}

	print(utils.GetMD5Hash(fmt.Sprint(jsonArray[0])))
	print(utils.GetMD5Hash(fmt.Sprint(jsonArray[1])))

	expectedResult := fmt.Sprint(map[string]bool{
		utils.GetMD5Hash(fmt.Sprint(jsonArray[0])): true,
		utils.GetMD5Hash(fmt.Sprint(jsonArray[1])): true,
	})

	testCases := []struct {
		name           string
		filepath       string
		expectedResult string
		expectedError  error
	}{
		{
			name:           "file not found",
			filepath:       "",
			expectedResult: "",
			expectedError:  &fs.PathError{},
		},
		{
			name:           "syntax error in file beginning",
			filepath:       testFilepathMalformedBegin,
			expectedResult: "",
			expectedError:  &json.SyntaxError{},
		},
		{
			name:           "cannot decode object",
			filepath:       testFilepathInvalidObject,
			expectedResult: "",
			expectedError:  &json.SyntaxError{},
		},
		{
			name:           "syntax error in file ending",
			filepath:       testFilepathMalformedEnd,
			expectedResult: "",
			expectedError:  &json.SyntaxError{},
		},
		{
			name:           "file processed successfully",
			filepath:       testFilepathA,
			expectedResult: expectedResult,
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processResult := Process(tc.filepath)

			if processResult.Result != tc.expectedResult {
				t.Fatal(utils.GetTestFailMessage(target, tc.expectedResult, processResult.Result))
			}

			if utils.AssertErrors(tc.expectedError, processResult.Err) {
				t.Fatal(utils.GetTestFailMessage(target, tc.expectedError, processResult.Err))
			}
		})
	}
}

func TestEqualSize(t *testing.T) {
	target := "EqualSize"

	testCases := []struct {
		name           string
		filepathA      string
		filepathB      string
		expectedResult bool
		expectedError  error
	}{
		{
			name:           "files have the same size",
			filepathA:      testFilepathA,
			filepathB:      testFilepathB,
			expectedResult: true,
			expectedError:  nil,
		},
		{
			name:           "files don't have the same size",
			filepathA:      testFilepathA,
			filepathB:      testFilepathC,
			expectedResult: false,
			expectedError:  nil,
		},
		{
			name:           "file A not found",
			filepathA:      "",
			filepathB:      testFilepathB,
			expectedResult: false,
			expectedError:  &fs.PathError{},
		},
		{
			name:           "file B not found",
			filepathA:      testFilepathA,
			filepathB:      "",
			expectedResult: false,
			expectedError:  &fs.PathError{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			equal, err := EqualSize(tc.filepathA, tc.filepathB)

			if equal != tc.expectedResult {
				t.Fatal(utils.GetTestFailMessage(target, tc.expectedResult, equal))
			}

			if utils.AssertErrors(tc.expectedError, err) {
				t.Fatal(utils.GetTestFailMessage(target, tc.expectedError, err))
			}
		})
	}
}

func TestGetFileSize(t *testing.T) {
	target := "GetFileSize"

	testCases := []struct {
		name          string
		filepath      string
		expectedSize  int64
		expectedError error
	}{
		{
			name:          "empty filepath",
			filepath:      "",
			expectedSize:  -1,
			expectedError: &fs.PathError{},
		},
		{
			name:          "file not found",
			filepath:      "invalid_filepath",
			expectedSize:  -1,
			expectedError: &fs.PathError{},
		},
		{
			name:          "file exists",
			filepath:      testFilepathA,
			expectedSize:  196,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size, err := getFileSize(tc.filepath)

			if size != tc.expectedSize {
				t.Fatal(utils.GetTestFailMessage(target, tc.expectedSize, size))
			}

			if utils.AssertErrors(tc.expectedError, err) {
				t.Fatal(utils.GetTestFailMessage(target, tc.expectedError, err))
			}
		})
	}
}
