package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotFound          = errors.New("file not found")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	src, err := os.OpenFile(fromPath, os.O_RDONLY, os.FileMode(0o755))
	if err != nil {
		return ErrFileNotFound
	}
	defer src.Close()

	srcInfo, err := src.Stat()
	switch {
	case err != nil:
		return ErrUnsupportedFile
	case srcInfo.IsDir():
		return ErrFileNotFound
	case offset > srcInfo.Size():
		return ErrOffsetExceedsFileSize
	}

	if offset > 0 {
		_, err = src.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
	}

	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	dstLen := srcInfo.Size()
	if dstLen > limit {
		dstLen = limit
	}

	bar := pb.Start64(dstLen)
	defer bar.Finish()
	barSrc := bar.NewProxyReader(src)

	if limit == 0 {
		_, err = io.Copy(dst, barSrc)
	} else {
		_, err = io.CopyN(dst, barSrc, limit)
	}
	if err != nil && err != io.EOF { //nolint:errorlint
		return err
	}

	return nil
}
