package xlogrus

import (
	"github.com/pkg/errors"
	"io"
	"os"
)

func NewFileHandler(fileName string) (io.WriteCloser, error)  {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return file, nil
}
