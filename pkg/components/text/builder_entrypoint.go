package text

import "ChromaTerm/pkg/ansi/color"

// New returns a new text builder.
func New() *Builder {
	return &Builder{}
}

// New returns a new text builder with text.
func Text(text string) *Builder {
	return &Builder{text: text}
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
func Foreground4Bit(c color.FG4Bit) *Builder {
	return New().Foreground4Bit(c)
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
func Background4Bit(c color.BG4Bit) *Builder {
	return New().Background4Bit(c)
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
func Color4Bit(fg color.FG4Bit, bg color.BG4Bit) *Builder {
	return New().Color4Bit(fg, bg)
}

// Set 8-bit foreground color. See [ChromaTerm/pkg/ansi/color/8b.md].
/*
 * 0-  7:   standard colors (as in 4-bit colors 30–37)
 * 8- 15:   high intensity colors (as in 4-bit colors 90–97)
 * 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
 * 232-255: gray-scale from dark to light in 24 steps
 */
func Foreground8Bit(n byte) *Builder {
	return New().Foreground8Bit(n)
}

// Set 8-bit background color. See [ChromaTerm/pkg/ansi/color/8b.md].
/*
 * 0-  7:   standard colors (as in 4-bit colors 30–37)
 * 8- 15:   high intensity colors (as in 4-bit colors 90–97)
 * 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
 * 232-255: gray-scale from dark to light in 24 steps
 */
func Background8Bit(n byte) *Builder {
	return New().Background8Bit(n)
}

// Set 24-bit foreground color with HEX (000 | #000 | 000000 | #000000).
func Foreground24BitHex(hex string) *Builder {
	return New().Foreground24BitHex(hex)
}

// Set 24-bit background color with HEX (000 | #000 | 000000 | #000000).
func Background24BitHex(hex string) *Builder {
	return New().Background24BitHex(hex)
}

// Set 24-bit foreground color with RGB. See [ChromaTerm/pkg/ansi/color/24b.md].
func Foreground24Bit(r, g, b byte) *Builder {
	return New().Foreground24Bit(r, g, b)
}

// Set 24-bit background color with RGB. See [ChromaTerm/pkg/ansi/color/24b.md].
func Background24Bit(r, g, b byte) *Builder {
	return New().Background24Bit(r, g, b)
}

// Bold or increased intensity.
func Bold() *Builder {
	return New().Bold()
}

// Faint (decreased intensity).
func Faint() *Builder {
	return New().Faint()
}

// Italic (Not widely supported, may be treated as inverse or blink).
func Italic() *Builder {
	return New().Italic()
}

// Underline.
func Underline() *Builder {
	return New().Underline()
}

// Crossed-out, or strike: Characters legible but marked as if for deletion.
// Not supported in Terminal.app.
func Strike() *Builder {
	return New().Strike()
}

// Doubly underlined (not bold - disables bold intensity on several terminals).
func DoubleUnderline() *Builder {
	return New().DoubleUnderline()
}
