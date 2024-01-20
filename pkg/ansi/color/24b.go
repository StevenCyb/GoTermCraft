package color

import (
	"ChromaTerm/pkg/ansi"
)

// Set 24-bit foreground color with RGB.
func Foreground24Bit(r, g, b byte) string {
	return "\033[38;2;" + ansi.B2S(r) + ";" + ansi.B2S(g) + ";" + ansi.B2S(b) + "m"
}

// Set 24-bit background color with RGB.
func Background24Bit(r, g, b byte) string {
	return "\033[48;2;" + ansi.B2S(r) + ";" + ansi.B2S(g) + ";" + ansi.B2S(b) + "m"
}

// Set 24-bit foreground color with HEX (000 | #000 | 000000 | #000000).
func Foreground24BitHex(hex string) string {
	r, g, b := hex2rgb(hex)

	return "\033[38;2;" + ansi.B2S(r) + ";" + ansi.B2S(g) + ";" + ansi.B2S(b) + "m"
}

// Set 24-bit background color with HEX (000 | #000 | 000000 | #000000).
func Background24BitHex(hex string) string {
	r, g, b := hex2rgb(hex)

	return "\033[48;2;" + ansi.B2S(r) + ";" + ansi.B2S(g) + ";" + ansi.B2S(b) + "m"
}

func hex2rgb(hex string) (r, g, b byte) {
	if hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) > 6 {
		hex = hex[:6]
	} else if len(hex) > 3 && len(hex) < 6 {
		hex = hex[:3]
	}

	rgb := make([]byte, 3)

	if len(hex) == 3 {
		rgb[0] = ansi.HexS2B(string(hex[0]) + string(hex[0]))
		rgb[1] = ansi.HexS2B(string(hex[1]) + string(hex[1]))
		rgb[2] = ansi.HexS2B(string(hex[2]) + string(hex[2]))
	} else {
		rgb[0] = ansi.HexS2B(string(hex[0]) + string(hex[1]))
		rgb[1] = ansi.HexS2B(string(hex[2]) + string(hex[3]))
		rgb[2] = ansi.HexS2B(string(hex[4]) + string(hex[5]))
	}

	return rgb[0], rgb[1], rgb[2]
}
