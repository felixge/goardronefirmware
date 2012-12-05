package navdata

import (
	"encoding/binary"
	"bytes"
	"io"
	"os"
)

type Reader struct{
	file *os.File
}

func NewReader() (*Reader, error) {
	// TODO set baudrate and other options (right now we get away without because
	// the official firmware configures the device before we kill it)
	file, err := os.OpenFile("/dev/ttyO1", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	reader := &Reader{}
	reader.file = file
	return reader, nil
}

// TODO figure out the correct structure coming from the tty (seems like
// AccTemp / GyroTemp are related to Ultrasound.
type Navdata struct{
	Seq uint16
	Acc [3]uint16
	Gyro [3]uint16
	Gyro110 [2]uint16
	AccTemp uint16
	GyroTemp uint16
	Ultrasound uint16
}

// ReadNavdata reads a single Navdata struct from the navdata serialport.
func (r *Reader) ReadNavdata() (*Navdata, error) {
	var size uint16
	for {
		err := binary.Read(r.file, binary.LittleEndian, &size)
		if err != nil {
			return nil, err
		}
		if size == 58 {
			break
		}
	}

	raw := make([]byte, size)
	_, err := io.ReadAtLeast(r.file, raw, int(size))
	if err != nil {
		return nil, err
	}

	data := &Navdata{}
	reader := bytes.NewReader(raw)

	err = binary.Read(reader, binary.LittleEndian, data)
	if err != nil{
		return nil, err
	}

	return data, nil
}
