package file

import (
	"io/fs"
	"os"
)

func GetFileSize(filepath string) (int64, error) {
	var fStats fs.FileInfo
	var err error

	if fStats, err = os.Stat(filepath); err != nil {
		return -1, err
	}

	return fStats.Size(), nil
}
