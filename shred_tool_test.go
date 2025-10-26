package shred

import (
	"os"
	"path/filepath"
	"testing"
)

func TestShred(t *testing.T) {
	const filename = "testFile"
	const data = "AAAAAAAAAAAAAAAAAAAA"

	filePath := filepath.Join(os.TempDir(), filename)
	file, err := os.Create(filePath)
	file.Write([]byte(data))

	if err != nil {
		t.Fatal("Failed to create testFile", err)
	}

	file.Close()

	if Shred(filePath) != nil {
		t.Error(err)
	}

	_, err = os.Open(filePath)
	if err == nil {
		t.Error("Failed to delete file", filePath)
	}

}

func TestRandomWrite(t *testing.T) {
	const filename = "testFile"
	const data = "AAAAAAAAAAAAAAAAAAAA"

	filePath := filepath.Join(os.TempDir(), filename)
	file, err := os.Create(filePath)
	file.Write([]byte(data))
	if err != nil {
		t.Fatal("Failed to create testFile", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatal("Failed to obtain file info", err)
	}

	fileSize := fileInfo.Size()

	numBytesWritten, err := randomWrite(file)

	if err != nil {
		t.Error(err)
	}

	if int64(numBytesWritten) != fileSize {
		t.Error("Number of random bytes written", numBytesWritten, "differs from fileSize", fileSize)
	}

}
