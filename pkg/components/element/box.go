package element

type BoxTitleLayout byte

const (
	BoxTitleLayoutTopLeft BoxTitleLayout = iota
	BoxTitleLayoutTopCenter
	BoxTitleLayoutTopRight
)

type Box struct {
	Element
	title       string
	titleLayout BoxTitleLayout
	border      bool
	borderStyle BorderStyle
}

func NewBox(width, height byte) *Box {
	if width < 4 {
		width = 4
	}
	if height < 4 {
		height = 4
	}

	box := &Box{
		Element: Element{
			width:  width,
			height: height,
		},
		borderStyle: DefaultBorderStyle(),
	}

	box.Element.refreshFunc = box.refresh

	return box
}

func (b *Box) refresh() {
	maxY := b.y + b.height
	maxX := b.x + b.width
	canvasElement := make([][]*rune, maxY)

	for y := byte(b.y); y < maxY; y++ {
		canvasElement[y] = make([]*rune, maxX)

		for x := byte(b.x); x < maxX; x++ {
			switch {
			case y == b.y && x == b.x:
				canvasElement[y][x] = &b.borderStyle.TopLeft
			case y == b.y && x == maxX-1:
				canvasElement[y][x] = &b.borderStyle.TopRight
			case y == maxY-1 && x == b.x:
				canvasElement[y][x] = &b.borderStyle.BottomLeft
			case y == maxY-1 && x == maxX-1:
				canvasElement[y][x] = &b.borderStyle.BottomRight
			case y == b.y:
				switch b.titleLayout {
				case BoxTitleLayoutTopLeft:
					if x-b.x-1 < byte(len(b.title)) {
						char := rune(b.title[x-b.x-1])
						canvasElement[y][x] = &char

						continue
					}
				case BoxTitleLayoutTopCenter:
					tLen := byte(len(b.title))
					center := byte(float32(maxX-b.x) / 2)
					left := b.x + (center - byte(tLen/2))
					if x >= left &&
						x < left+tLen-1 {
						char := rune(b.title[x-left])
						canvasElement[y][x] = &char

						continue
					}
				case BoxTitleLayoutTopRight:
					if x >= maxX-1-byte(len(b.title)) {
						char := rune(b.title[x-(maxX-1-byte(len(b.title)))])
						canvasElement[y][x] = &char

						continue
					}
				}
				canvasElement[y][x] = &b.borderStyle.Top
			case y == maxY-1:
				canvasElement[y][x] = &b.borderStyle.Bottom
			case x == b.x:
				canvasElement[y][x] = &b.borderStyle.Left
			case x == maxX-1:
				canvasElement[y][x] = &b.borderStyle.Right
			}
		}
	}

	b.copyToCanvas(canvasElement)
}

func (b *Box) SetTitle(title string, layout BoxTitleLayout) {
	if len(title) > 255 {
		title = title[:255]
	}

	b.title = title
	b.titleLayout = layout
}

func (b *Box) UseBorder() {
	b.border = true
}

func (b *Box) SetBorderStyle(style BorderStyle) {
	b.borderStyle = style
}

// ┌────────────────────────────────────── Wgservices(all)[2] ───────────────────────────────────────┐
// │ NAMESPACE↑              NAME                                               AGE                  │
// │ default                 cc3a93bc-e72b-4242-8906-f3a322dbebf6               2d22h                │
// │ default                 cf9a0e16-5558-422a-8a52-87bf46c1d38c               2d22h                │
// │                                                                                                 │
// │                                                                                                 │
// │                                                                                                 │
// │                                                                                                 │
// │                                                                                                 │
// │                                                                                                 │
// │                                                                                                 │
// │                                                                                                 │
// │                                                                                                 │
// │                                                                                                 │
// └─────────────────────────────────────────────────────────────────────────────────────────────────┘
