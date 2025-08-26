package rawmode

import (
	"os"
)

var (
	// GetMode from file descriptor
	GetMode func(std *os.File) (any, error)
	// SetMode to file descriptor
	SetMode func(std *os.File, mode any) error
	// SetRawMode to file descriptor enriching existign mode with raw console flags
	SetRawMode func(std *os.File, mode any) error
	// Read from file descriptor to buffer
	Read func(std *os.File, buf []byte) (int, error)

	TCSETS uintptr
	TCGETS uintptr
)
