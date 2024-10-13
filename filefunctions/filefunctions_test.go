package filefunctions

import "testing"

func TestOpenFileCloseFile(_ *testing.T) {
	f := OpenFile("filefunctions.go")
	CloseFile(f)
}
