package navdata

import (
	"os"
)

type Reader struct{
	*os.File
}

func NewReader() (*Reader, error) {
	file, err := os.OpenFile("/dev/ttyO1", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	reader := &Reader{file}
	return reader, nil
}
