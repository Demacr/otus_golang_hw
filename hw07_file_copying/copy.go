package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

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

	_, _ = fromFile.Seek(offset, io.SeekStart)

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

	_, err = io.CopyN(toFile, fromFile, limit)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
