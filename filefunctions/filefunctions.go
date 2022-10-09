package filefunctions

import (
	"os"
)

// Check checks error values
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// OpenFile returns a file pointer. The file path is relative. CloseFile must be called on the returned file pointer.
func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	Check(err)
	return file
}

// CloseFile closes a file pointer.
func CloseFile(f *os.File) {
	err := f.Close()
	Check(err)
}
