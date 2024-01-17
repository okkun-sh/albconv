package cmd

import (
	"albconv/cmd/conv"
	"bytes"
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
)

type indexedEntry struct {
	index int
	data  []byte
}

type convError struct {
	filepath string
	err      error
}

func Run(filepaths []string) (string, error) {
	if len(filepaths) == 0 {
		return "", errors.New("no files specified")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	entriesChan := make(chan indexedEntry, len(filepaths))
	errChan := make(chan convError, 1)

	jc := conv.NewJSONConverter()
	for i, filepath := range filepaths {
		wg.Add(1)
		go func(i int, filepath string) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				d, err := Decompress(filepath)
				if err != nil {
					errChan <- convError{filepath: filepath, err: err}
					cancel()
					return
				}
				entries, err := jc.Convert(d)
				if err != nil {
					errChan <- convError{filepath: filepath, err: err}
					cancel()
					return
				}
				entriesChan <- indexedEntry{index: i, data: entries}
			}
		}(i, filepath)
	}

	wg.Wait()
	close(entriesChan)

	var entries []indexedEntry
	for e := range entriesChan {
		entries = append(entries, e)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].index < entries[j].index
	})

	var data []byte
	data = append(data, '[')
	for i, e := range entries {
		if i > 0 {
			data = append(data, ',')
		}
		tData := bytes.Trim(e.data, "[]")
		data = append(data, tData...)
	}
	data = append(data, ']')

	select {
	case convErr := <-errChan:
		return "", fmt.Errorf("error processing file %s: %v", convErr.filepath, convErr.err)
	default:
		return string(data), nil
	}
}
