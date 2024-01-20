package element

import (
	"os"

	"golang.org/x/term"
)

type Engine struct {
	elements []ElementI
	canvas   [][]rune
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) terminalSize() (int, int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0
	}

	return w, h
}

func (e *Engine) copyCanvas(canvas [][]*rune) {
	for len(canvas) > len(e.canvas) {
		e.canvas = append(e.canvas, make([]rune, 0))
	}

	for x := 0; x < len(canvas); x++ {
		for len(canvas[x]) > len(e.canvas[x]) {
			e.canvas[x] = append(e.canvas[x], ' ')
		}

		for y := 0; y < len(canvas[x]); y++ {
			if canvas[x][y] == nil {
				continue
			}

			e.canvas[x][y] = *canvas[x][y]
		}
	}
}

func (e *Engine) Refresh() {
	e.canvas = make([][]rune, 0)

	for _, elm := range e.elements {
		elm.Refresh()
	}

	w, h := e.terminalSize()

	for y := 0; y < h; y++ {
		if y >= len(e.canvas) {
			break
		}

		for x := 0; x < w; x++ {
			if x >= len(e.canvas[y]) {
				break
			}

			print(string(e.canvas[y][x]))
		}

		print("\n")
	}
}

func (e *Engine) AddElement(elm ElementI) {
	elm.copyToCanvasFunc(e.copyCanvas)

	e.elements = append(e.elements, elm)
}
