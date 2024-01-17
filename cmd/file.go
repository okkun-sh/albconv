package cmd

import (
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"os"
)

func Decompress(filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 512)
	if _, err = f.Read(buf); err != nil {
		return nil, err
	}

	cType := http.DetectContentType(buf)
	if cType == "application/x-gzip" {
		if _, err = f.Seek(0, 0); err != nil {
			return nil, err
		}
		gz, err := gzip.NewReader(f)
		if err != nil {
			return nil, err
		}
		defer gz.Close()

		d, err := io.ReadAll(gz)
		if err != nil {
			return nil, err
		}
		return d, nil
	}
	return nil, errors.New("file type is not gzip: " + cType)
}
