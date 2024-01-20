package text

import (
	"ChromaTerm/pkg/ansi/color"
	"ChromaTerm/pkg/ansi/csi"
	"ChromaTerm/pkg/ansi/sgr"
)

// Builder is a text builder.
type Builder struct {
	color      string
	decoration string
	text       string
}

// Returns a string that with all attributes.
func (t *Builder) String() string {
	return t.color + t.decoration + t.text
}

// Returns a string that with all attributes.
// Will automatically reset include `sgr.Reset()`.
func (t *Builder) StringWithReset() string {
	return t.color + t.decoration + t.text + sgr.Reset()
}

// Set text that interpret all previous attributes.
func (t *Builder) Text(text string) *Builder {
	t.text = text

	return t
}

// Set foreground color with 8-bit. See [ChromaTerm/pkg/ansi/color/8b.md].
/*
 * FG | COLOR
 * 30 | Black
 * 31 | Red
 * 32 | Green
 * 33 | Yellow
 * 34 | Blue
 * 35 | Magenta
 * 36 | Cyan
 * 37 | White
 * 90 | Bright Black (Gray)
 * 91 | Bright Red
 * 92 | Bright Green
 * 93 | Bright Yellow
 * 94 | Bright Blue
 * 95 | Bright Magenta
 * 96 | Bright Cyan
 * 97 | Bright White
 */
func (t *Builder) Foreground4Bit(c color.FG4Bit) *Builder {
	t.color = color.Foreground4Bit(c)

	return t
}

// Set background color with 4-bit. See [ChromaTerm/pkg/ansi/color/8b.md].
/*
 * BG  | COLOR
 * 40  | Black
 * 41  | Red
 * 42  | Green
 * 43  | Yellow
 * 44  | Blue
 * 45  | Magenta
 * 46  | Cyan
 * 47  | White
 * 100 | Bright Black (Gray)
 * 101 | Bright Red
 * 102 | Bright Green
 * 103 | Bright Yellow
 * 104 | Bright Blue
 * 105 | Bright Magenta
 * 106 | Bright Cyan
 * 107 | Bright White
 */
func (t *Builder) Background4Bit(c color.BG4Bit) *Builder {
	t.color = color.Background4Bit(c)

	return t
}

// Set foreground and background color with 4-bit. See [ChromaTerm/pkg/ansi/color/8b.md].
/*
 * FG | BG  | COLOR
 * 30 | 40  | Black
 * 31 | 41  | Red
 * 32 | 42  | Green
 * 33 | 43  | Yellow
 * 34 | 44  | Blue
 * 35 | 45  | Magenta
 * 36 | 46  | Cyan
 * 37 | 47  | White
 * 90 | 100 | Bright Black (Gray)
 * 91 | 101 | Bright Red
 * 92 | 102 | Bright Green
 * 93 | 103 | Bright Yellow
 * 94 | 104 | Bright Blue
 * 95 | 105 | Bright Magenta
 * 96 | 106 | Bright Cyan
 * 97 | 107 | Bright White
 */
func (t *Builder) Color4Bit(fg color.FG4Bit, bg color.BG4Bit) *Builder {
	t.color = color.Color4Bit(fg, bg)

	return t
}

// Set 8-bit foreground color. See [ChromaTerm/pkg/ansi/color/8b.md].
/*
 * 0-  7:   standard colors (as in 4-bit colors 30–37)
 * 8- 15:   high intensity colors (as in 4-bit colors 90–97)
 * 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
 * 232-255: gray-scale from dark to light in 24 steps
 */
func (t *Builder) Foreground8Bit(n byte) *Builder {
	t.color = color.Foreground8Bit(n)

	return t
}

// Set 8-bit background color. See [ChromaTerm/pkg/ansi/color/8b.md].
/*
 * 0-  7:   standard colors (as in 4-bit colors 30–37)
 * 8- 15:   high intensity colors (as in 4-bit colors 90–97)
 * 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
 * 232-255: gray-scale from dark to light in 24 steps
 */
func (t *Builder) Background8Bit(n byte) *Builder {
	t.color = color.Background8Bit(n)

	return t
}

// Set 24-bit foreground color with RGB. See [ChromaTerm/pkg/ansi/color/24b.md].
func (t *Builder) Foreground24Bit(r, g, b byte) *Builder {
	t.color = color.Foreground24Bit(r, g, b)

	return t
}

// Set 24-bit background color with RGB. See [ChromaTerm/pkg/ansi/color/24b.md].
func (t *Builder) Background24Bit(r, g, b byte) *Builder {
	t.color = color.Background24Bit(r, g, b)

	return t
}

// Set 24-bit foreground color with HEX (000 | #000 | 000000 | #000000).
func (t *Builder) Foreground24BitHex(hex string) *Builder {
	t.color = color.Foreground24BitHex(hex)

	return t
}

// Set 24-bit background color with HEX (000 | #000 | 000000 | #000000).
func (t *Builder) Background24BitHex(hex string) *Builder {
	t.color = color.Background24BitHex(hex)

	return t
}

// Default foreground color.
func (t *Builder) DefaultForegroundColor() *Builder {
	t.color = color.DefaultForegroundColor()

	return t
}

// Default foreground color.
func (t *Builder) DefaultBackgroundColor() *Builder {
	t.color = color.DefaultBackgroundColor()

	return t
}

// Default background color.
func (t *Builder) Bold() *Builder {
	t.decoration = csi.SelectGraphicRendition(1)

	return t
}

// Faint (decreased intensity).
func (t *Builder) Faint() *Builder {
	t.decoration = csi.SelectGraphicRendition(2)

	return t
}

// Italic (Not widely supported, may be treated as inverse or blink).
func (t *Builder) Italic() *Builder {
	t.decoration = csi.SelectGraphicRendition(3)

	return t
}

// Underline.
func (t *Builder) Underline() *Builder {
	t.decoration = csi.SelectGraphicRendition(4)

	return t
}

// Crossed-out, or strike: Characters legible but marked as if for deletion.
// Not supported in Terminal.app.
func (t *Builder) Strike() *Builder {
	t.decoration = csi.SelectGraphicRendition(9)

	return t
}

// Doubly underlined (not bold - disables bold intensity on several terminals).
func (t *Builder) DoubleUnderline() *Builder {
	t.decoration = csi.SelectGraphicRendition(21)

	return t
}
