package file

import (
	"io/fs"
	"testing"

	"github.com/joaopires/hatch/internal/utils"
)

const testFilepathA = "../../resources/equal/unordered_small_file_before_backup.json"
const testFilepathB = "../../resources/equal/unordered_small_file_after_backup.json"
const testFilepathC = "../../resources/equal/unordered_medium_file_before_backup.json"

func TestProcess(t *testing.T) {

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
			name:          "file does not exists",
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
