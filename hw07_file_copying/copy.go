package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Copy(fromPath string, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	fileFromStat, err := fromFile.Stat()
	if err != nil {
		return err
	}
	if !fileFromStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if fileFromStat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer toFile.Close()

	_, _ = fromFile.Seek(offset, 0)

	// Handle limit == 0
	if limit == 0 || fileFromStat.Size()-offset < limit {
		limit = fileFromStat.Size() - offset
	}

	// Start progressbar
	progressBar := pb.New64(limit)
	progressBar.SetUnits(pb.U_BYTES)
	progressBar.Start()
	defer progressBar.Finish()

	// Copying
	var position int64
	var N int64 = 4096
	buffer := make([]byte, N)
	for position < limit {
		var read int
		var errRead error

		read, errRead = fromFile.Read(buffer[:min(N, limit-position)])

		position += int64(read)
		progressBar.Set64(position)

		_, errWrite := toFile.Write(buffer[:read])
		if errWrite != nil {
			return err
		}

		if errRead == io.EOF {
			break
		}
	}
	return nil
}
