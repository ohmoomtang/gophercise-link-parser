package utils

import (
	"io"
	"os"
)

func ReadHTMLFile(filename string) (io.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
