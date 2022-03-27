package filefunctions

import (
	"os"
	"path/filepath"
)

// Check checks error values
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// OpenFile return a file pointer. The file path is relative to the current working directory. CloseFile must be called on the returned file pointer.
func OpenFile(path string) *os.File {
	absPath, err := filepath.Abs(path)
	Check(err)
	file, err := os.Open(absPath)
	Check(err)
	return file
}

// CloseFile closes a file pointer.
func CloseFile(f *os.File) {
	err := f.Close()
	Check(err)
}
