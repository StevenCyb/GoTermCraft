package form

import (
	"ChromaTerm/pkg/ansi/csi"
	"ChromaTerm/pkg/terminal"
	"os"
	"strings"
)

type input struct {
	terminal   *terminal.Terminal
	isSecret   bool
	label      string
	onInput    func(string)
	onValidate func(string) bool
}

func NewInput() *input {
	terminal, _ := terminal.New(os.Stdin, os.Stdout)

	return &input{
		label:    "Input:",
		terminal: terminal,
	}
}

func (i *input) Terminal(terminal *terminal.Terminal) *input {
	i.terminal = terminal

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

	for {
		for {
			char := i.terminal.ReadByte()
			if char == nil {
				continue
			}

			if *char == 13 {
				break
			}

			if i.isSecret {
				i.terminal.Print(csi.CursorBack(byte(len(input))) + strings.Repeat("*", len(input)))
			} else {
				i.terminal.Print(*char)
			}

			input += string(*char)

			if i.onInput != nil {
				i.onInput(input)
			}
		}

		if i.onValidate != nil && !i.onValidate(input) {
			i.terminal.Print(csi.CursorPreviousLine(1) + csi.EraseInLine(csi.EraseLineEntire))

			input = ""

			continue
		} else if hideInput {
			i.terminal.Print(csi.CursorPreviousLine(1) + csi.EraseInLine(csi.EraseLineEntire))
		}

		break
	}

	return &input, nil
}
