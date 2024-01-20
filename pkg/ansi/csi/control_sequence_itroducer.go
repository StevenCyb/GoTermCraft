package csi

import "ChromaTerm/pkg/ansi"

// Moves the cursor n cells in the given direction.
// If the cursor is already at the edge of the screen, this has no effect.
func CursorUp(n byte) string {
	if n <= 1 {
		return "\033[A"
	}

	return "\033[" + ansi.B2S(n) + "A"
}

// Moves the cursor n cells in the given direction.
// If the cursor is already at the edge of the screen, this has no effect.
func CursorDown(n byte) string {
	if n <= 1 {
		return "\033[B"
	}

	return "\033[" + ansi.B2S(n) + "B"
}

// Moves the cursor n cells in the given direction.
// If the cursor is already at the edge of the screen, this has no effect.
func CursorForward(n byte) string {
	if n <= 1 {
		return "\033[C"
	}

	return "\033[" + ansi.B2S(n) + "C"
}

// Moves the cursor n cells in the given direction.
// If the cursor is already at the edge of the screen, this has no effect.
func CursorBack(n byte) string {
	if n <= 1 {
		return "\033[D"
	}

	return "\033[" + ansi.B2S(n) + "D"
}

// Moves cursor to beginning of the line n lines down.
func CursorNextLine(n byte) string {
	if n <= 1 {
		return "\033[E"
	}

	return "\033[" + ansi.B2S(n) + "E"
}

// Moves cursor to beginning of the line n lines up.
func CursorPreviousLine(n byte) string {
	if n <= 1 {
		return "\033[F"
	}

	return "\033[" + ansi.B2S(n) + "F"
}

// Moves the cursor to column n.
func CursorHorizontalAbsolute(n byte) string {
	if n <= 1 {
		return "\033[G"
	}

	return "\033[" + ansi.B2S(n) + "G"
}

// Moves the cursor to row/column. The values are 1-based, and
// 1/1 corresponded to top left corner.
func CursorPosition(row, col byte) string {
	if row == 0 {
		row = 1
	}
	if col == 0 {
		col = 1
	}

	return "\033[" + ansi.B2S(row) + ";" + ansi.B2S(col) + "H"
}

// Show cursor.
func ShowCursor() string {
	return "\033[?25h"
}

// Hide cursor.
func HideCursor() string {
	return "\033[?25l"
}

type EraseScreen byte

const (
	// Clear from cursor to end of screen.
	EraseScreenToEnd EraseScreen = iota
	// Clear from cursor to beginning of the screen.
	EraseScreenToBeginning
	// Clear entire screen.
	EraseScreenEntire
	// Clear entire screen and delete all lines saved in the scroll back buffer.
	EraseScreenAndScrollBackBuffer
)

// Clears part of the screen.
/*
 * EraseScreenToEnd(0)               = clear from cursor to end of screen
 * EraseScreenToBeginning(1)         = clear from cursor to beginning of the screen
 * EraseScreenEntire(2)              = clear entire screen
 * EraseScreenAndScrollBackBuffer(3) = clear entire screen and delete all lines saved in the scroll back buffer
 */
func EraseInScreen(n EraseScreen) string {
	if n > 3 {
		n = 3
	}

	return "\033[" + ansi.B2S(byte(n)) + "J"
}

type EraseLine byte

const (
	// Clear from cursor to end of line.
	EraseLineToEnd EraseLine = iota
	// Clear from cursor to beginning of the line.
	EraseLineToBeginning
	// Clear entire line.
	EraseLineEntire
)

// Erases part of the line from the cursor perspective.
// Cursor position does not change.
func EraseInLine(n EraseLine) string {
	if n > 2 {
		n = 2
	}

	return "\033[" + ansi.B2S(byte(n)) + "K"
}

// Scroll whole page up by lines. 
// New lines are added at the bottom.
func ScrollUp(n byte) string {
	if n <= 1 {
		return "\033[S"
	}

	return "\033[" + ansi.B2S(n) + "S"
}

// Scroll whole page down by lines.
// New lines are added at the top.
func ScrollDown(n byte) string {
	if n <= 1 {
		return "\033[T"
	}

	return "\033[" + ansi.B2S(n) + "T"
}

// Set Graphics Rendition is used to set text attributes, 
// such as bold, underline, and colors.
// Use SGR functions to set text attributes.
func SelectGraphicRendition(n byte) string {
	return "\033[" + ansi.B2S(n) + "m"
}
