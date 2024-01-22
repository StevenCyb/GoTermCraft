package terminal

import (
	"bufio"
	"fmt"
	"io"
	"strings"
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
	stdin      Reader
	stdout     Writer
	fd         uintptr
	readBuffer []byte
}

// New creates a new Terminal instance with the specified stdin and stdout.
func New(stdin Reader, stdout Writer) (*Terminal, error) {
	terminal := &Terminal{
		fd:         stdout.Fd(),
		stdin:      stdin,
		stdout:     stdout,
		readBuffer: []byte{},
	}

	_, err := term.MakeRaw(terminal.fd)
	if err != nil {
		return nil, err
	}

	go func() {
		reader := bufio.NewReader(terminal.stdin)

		for {
			char, err := reader.ReadByte()
			if err != nil {
				continue
			}

			if char == 3 {
				terminal.Print("^C\n")
				syscall.Exit(0)
			} else if char == 13 {
				terminal.Print("\r\n")
			} else if char == 127 {
				terminal.Print("\b \b")
			}

			terminal.readBuffer = append(terminal.readBuffer, char)
		}
	}()

	return terminal, nil
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
func (t *Terminal) ReadByte() *byte {
	if len(t.readBuffer) == 0 {
		return nil
	}

	char := t.readBuffer[0]
	t.readBuffer = t.readBuffer[1:]

	return &char
}

// Print prints a string to the terminal.
func (t *Terminal) Print(a ...any) {
	out := strings.ReplaceAll(fmt.Sprint(a...), "\n", "\r\n")

	fmt.Fprint(t.stdout, out)
}

// Printf formats and prints to the terminal.
func (t *Terminal) Printf(format string, a ...any) {
	t.Print(fmt.Sprintf(format, a...))
}

// Println prints a line to the terminal.
func (t *Terminal) Println(a ...any) {
	if a == nil {
		return
	}

	a = append(a, "\n")

	t.Print(a...)
}
