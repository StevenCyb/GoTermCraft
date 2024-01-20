package color

import "ChromaTerm/pkg/ansi"

// Set 8-bit foreground color. See [ChromaTerm/pkg/ansi/color/8b.md].
/*
 * 0-  7:   standard colors (as in 4-bit colors 30–37)
 * 8- 15:   high intensity colors (as in 4-bit colors 90–97)
 * 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
 * 232-255: gray-scale from dark to light in 24 steps
 */
func Foreground8Bit(n byte) string {
	return "\033[38;5;" + ansi.B2S(n) + "m"
}

// Set 8-bit background color. See [ChromaTerm/pkg/ansi/color/8b.md].
/*
 * 0-  7:   standard colors (as in 4-bit colors 30–37)
 * 8- 15:   high intensity colors (as in 4-bit colors 90–97)
 * 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
 * 232-255: gray-scale from dark to light in 24 steps
 */
func Background8Bit(n byte) string {
	return "\033[48;5;" + ansi.B2S(n) + "m"
}
