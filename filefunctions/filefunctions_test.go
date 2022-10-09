package filefunctions

import "testing"

func TestOpenFileCloseFile(t *testing.T) {
	f := OpenFile("filefunctions.go")
	CloseFile(f)
}
