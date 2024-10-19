package filefunctions

import (
	"errors"
	"testing"
)

// TestOpenFile tests the OpenFile function.
func TestOpenFile(t *testing.T) {
	// Case 1: Open an existing file (e.g., "filefunctions.go").
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("OpenFile(\"filefunctions.go\") panicked unexpectedly: %v", r)
		}
	}()
	file := OpenFile("filefunctions.go")
	if file == nil {
		t.Errorf("Expected a valid *os.File, got nil")
	}

	// Close the file to ensure CloseFile works.
	CloseFile(file)

	// Case 2: Attempt to open a non-existent file, expecting a panic.
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("OpenFile(non-existent file) did not panic as expected")
		}
	}()
	OpenFile("non_existent_file.go")
}

// TestCloseFile tests the CloseFile function.
func TestCloseFile(t *testing.T) {
	// Open an existing file for testing CloseFile.
	file := OpenFile("filefunctions.go")

	// Case 1: Close an open file
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("CloseFile(filefunctions.go) panicked unexpectedly: %v", r)
		}
	}()
	CloseFile(file)

	// Case 2: Try to close an already closed file, expecting a panic.
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CloseFile(closed file) did not panic as expected")
		}
	}()
	CloseFile(file) // This should panic as the file is already closed.
}

// TestCheck tests the Check function.
func TestCheck(t *testing.T) {
	// Case 1: Check with nil error, should not panic.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Check(nil) panicked unexpectedly: %v", r)
		}
	}()
	Check(nil)

	// Case 2: Check with a non-nil error, should panic.
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Check(non-nil) did not panic as expected")
		}
	}()
	Check(errors.New("test error"))
}
