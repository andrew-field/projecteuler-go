package filefunctions_test

import (
	"testing"

	"github.com/andrew-field/projecteuler-go/filefunctions"
)

func TestOpenFileCloseFile(t *testing.T) {
	f := filefunctions.OpenFile("filefunctions.go")
	filefunctions.CloseFile(f)
}
