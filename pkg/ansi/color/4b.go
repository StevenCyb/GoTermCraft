package color

import "ChromaTerm/pkg/ansi"

type FG4Bit byte

type BG4Bit byte

const (
	ForegroundBlack FG4Bit = iota + 30
	ForegroundRed
	ForegroundGreen
	ForegroundYellow
	ForegroundBlue
	ForegroundMagenta
	ForegroundCyan
	ForegroundWhite
	ForegroundBrightBlack FG4Bit = iota + 82
	ForegroundBrightRed
	ForegroundBrightGreen
	ForegroundBrightYellow
	ForegroundBrightBlue
	ForegroundBrightMagenta
	ForegroundBrightCyan
	ForegroundBrightWhite
)

const (
	BackgroundBlack BG4Bit = iota + 40
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
	BackgroundBrightBlack BG4Bit = iota + 92
	BackgroundBrightRed
	BackgroundBrightGreen
	BackgroundBrightYellow
	BackgroundBrightBlue
	BackgroundBrightMagenta
	BackgroundBrightCyan
	BackgroundBrightWhite
)

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
func Foreground4Bit(n FG4Bit) string {
	if n < 30 || (n > 37 && n < 90) || n > 97 {
		n = ForegroundBlack
	}

	return "\033[" + ansi.B2S(byte(n)) + "m"
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
func Background4Bit(n BG4Bit) string {
	if n < 40 || (n > 47 && n < 100) || n > 107 {
		n = BackgroundBlack
	}

	return "\033[" + ansi.B2S(byte(n)) + "m"
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
func Color4Bit(fg FG4Bit, bg BG4Bit) string {
	if fg < 30 || (fg > 37 && fg < 90) || fg > 97 {
		fg = ForegroundBlack
	}
	if bg < 40 || (bg > 47 && bg < 100) || bg > 107 {
		bg = BackgroundBlack
	}

	return "\033[" + ansi.B2S(byte(fg)) + ";" + ansi.B2S(byte(bg)) + "m"
}
