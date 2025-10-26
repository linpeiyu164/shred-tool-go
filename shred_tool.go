package shred

import (
	"crypto/rand"
	"os"
)

func randomWrite(file *os.File) (n int, err error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	fileSize := fileInfo.Size()
	randomBytes := make([]byte, fileSize)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return 0, err
	}

	bytesWritten, err := file.Write(randomBytes)
	if err != nil {
		return 0, err
	}

	return bytesWritten, nil
}

func Shred(filePath string) (err error) {

	const iterations = 3

	file, err := os.OpenFile(filePath, os.O_WRONLY, 0700)

	if err != nil {
		return err
	}

	for i := 0; i < iterations; i++ {
		_, err := randomWrite(file)
		if err != nil {
			return err
		}
		file.Seek(0, 0)
	}

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
