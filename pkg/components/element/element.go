package element

type ElementI interface {
	copyToCanvasFunc(func(canvas [][]*rune))
	Refresh()
	Position(x, y byte)
	Resize(width, height byte)
	GetWidth() byte
	GetHeight() byte
	GetPosition() (x, y byte)
}

type Element struct {
	copyToCanvas func(canvas [][]*rune)
	refreshFunc  func()
	x, y         byte
	width        byte
	height       byte
}

func (e *Element) copyToCanvasFunc(copyToCanvasFunc func(canvas [][]*rune)) {
	e.copyToCanvas = copyToCanvasFunc
}

func (e *Element) Refresh() {
	if e.refreshFunc != nil {
		e.refreshFunc()
	}
}

func (e *Element) Position(x, y byte) {
	e.x = x
	e.y = y
}

func (e *Element) Resize(width, height byte) {
	e.width = width
	e.height = height
}

func (e *Element) GetWidth() byte {
	return e.width
}

func (e *Element) GetHeight() byte {
	return e.height
}

func (e *Element) GetPosition() (x, y byte) {
	return e.x, e.y
}
