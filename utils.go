package goutils

import (
	"io"
	"os"
)

// GetReader checks the CLI arguments provided to the program.
// If no argument is provided or if the first argument is "-", it returns os.Stdin.
// Otherwise, it attempts to open the file path provided in the first argument.
//
// The caller is responsible for calling .Close() on the returned io.ReadCloser.
func GetReader() (io.ReadCloser, error) {

	if len(os.Args) > 1 && os.Args[1] != "-" {
		file, err := os.Open(os.Args[1])
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return os.Stdin, nil
}
