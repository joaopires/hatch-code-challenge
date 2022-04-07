package file

import (
	"io/fs"
	"testing"

	"github.com/joaopires/hatch/internal/utils"
)

const testFilepath = "../../resources/json-before.json"

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
			filepath:      testFilepath,
			expectedSize:  196,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size, err := GetFileSize(tc.filepath)

			if size != tc.expectedSize {
				t.Fatal(utils.GetTestFailMessage(target, tc.expectedSize, size))
			}

			if err != nil {
				if tc.expectedError == nil {
					t.Fatal(utils.GetTestFailMessage(target, tc.expectedError, err))
				}

				if _, ok := err.(*fs.PathError); !ok {
					t.Fatal(utils.GetTestFailMessage(target, tc.expectedError, err))
				}
			}
		})
	}
}
