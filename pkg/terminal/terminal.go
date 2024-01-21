package terminal

import (
	"fmt"
	"io"
	"syscall"

	"github.com/moby/term"
)

type Reader interface {
	io.Reader
	Fd() uintptr
}

type Writer interface {
	io.Writer
	Fd() uintptr
}

type Terminal struct {
	stdin  Reader
	stdout Writer
	fd     uintptr
}

// New creates a new Terminal instance with the specified stdin and stdout.
func New(stdin Reader, stdout Writer) (*Terminal, error) {
	fd := stdout.Fd()

	// Set terminal to raw mode
	_, err := term.MakeRaw(fd)
	if err != nil {
		return nil, err
	}

	return &Terminal{
		fd:     fd,
		stdin:  stdin,
		stdout: stdout,
	}, nil
}

// Size returns the width and height of the terminal.
func (t *Terminal) Size() (int, int, error) {
	size, err := term.GetWinsize(t.fd)
	if err != nil {
		return 0, 0, err
	}

	return int(size.Width), int(size.Height), nil
}

// ReadByte reads a single byte from the terminal.
func (t *Terminal) ReadByte() (byte, error) {
	var b [1]byte
	_, err := syscall.Read(int(t.fd), b[:])
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

// Print prints a string to the terminal.
func (t *Terminal) Print(a ...any) {
	fmt.Fprint(t.stdout, a...)
}

// Printf formats and prints to the terminal.
func (t *Terminal) Printf(format string, a ...any) {
	fmt.Fprintf(t.stdout, format, a...)
}

// Println prints a line to the terminal.
func (t *Terminal) Println(a ...any) {
	fmt.Fprintln(t.stdout, a...)
}
