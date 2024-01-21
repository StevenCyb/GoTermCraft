package form

import (
	"ChromaTerm/pkg/ansi/csi"
	"bufio"
	"io"
	"os"
	"strings"
)

type input struct {
	stdout     io.Writer
	stdin      io.Reader
	isSecret   bool
	label      string
	onInput    func(string)
	onValidate func(string) bool
}

func NewInput() *input {
	return &input{
		label:  "Input:",
		stdin:  os.Stdin,
		stdout: os.Stdout,
	}
}

func (i *input) Stdout(stdout io.Writer) *input {
	i.stdout = stdout

	return i
}

func (i *input) Stdin(stdin io.Reader) *input {
	i.stdin = stdin

	return i
}

func (i *input) IsSecret() *input {
	i.isSecret = true

	return i
}

func (i *input) Label(label string) *input {
	i.label = label

	return i
}

func (i *input) OnValidate(onValidate func(string) bool) *input {
	i.onValidate = onValidate

	return i
}

func (i *input) OnInput(onInput func(string)) *input {
	i.onInput = onInput

	return i
}

func (i *input) Read(hideInput bool) (*string, error) {
	input := ""
	reader := bufio.NewReader(i.stdin)

	for {
		i.stdout.Write([]byte(i.label + ": "))

		for {
			char, err := reader.ReadByte()
			if err != nil {
				return nil, err
			}

			if char == '\n' {
				break
			}

			if i.isSecret {
				i.stdout.Write([]byte(csi.CursorBack(byte(len(input))) + strings.Repeat("*", len(input))))
			} else {
				i.stdout.Write([]byte{char})
			}

			input += string(char)

			if i.onInput != nil {
				i.onInput(input)
			}
		}

		if i.onValidate != nil && !i.onValidate(input) {
			i.stdout.Write([]byte(csi.CursorPreviousLine(1) + csi.EraseInLine(csi.EraseLineEntire)))

			input = ""

			continue
		} else if hideInput {
			i.stdout.Write([]byte(csi.CursorPreviousLine(1) + csi.EraseInLine(csi.EraseLineEntire)))
		}

		break
	}

	return &input, nil
}
