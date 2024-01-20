package sgr

import "ChromaTerm/pkg/ansi/csi"

// Reset all attributes.
func Reset() string {
	return csi.SelectGraphicRendition(0)
}

// Bold or increased intensity.
func Bold() string {
	return csi.SelectGraphicRendition(1)
}

// Faint (decreased intensity).
func Faint() string {
	return csi.SelectGraphicRendition(2)
}

// Italic (Not widely supported, may be treated as inverse or blink).
func Italic() string {
	return csi.SelectGraphicRendition(3)
}

// Underline.
func Underline() string {
	return csi.SelectGraphicRendition(4)
}

// Sets blinking to less than 150 n per minute.
func BlinkSlow() string {
	return csi.SelectGraphicRendition(5)
}

// Sets blinking to 150+ per minute; not widely supported.
func BlinkRapid() string {
	return csi.SelectGraphicRendition(6)
}

// Swap foreground and background colors; inconsistent emulation.
func ImageNegative() string {
	return csi.SelectGraphicRendition(7)
}

// Crossed-out, or strike: Characters legible but marked as if for deletion.
// Not supported in Terminal.app.
func Strike() string {
	return csi.SelectGraphicRendition(9)
}

// Doubly underlined (not bold - disables bold intensity on several terminals).
func DoubleUnderline() string {
	return csi.SelectGraphicRendition(21)
}

// Neither bold nor faint; color changes where intensity is implemented as such.
func NormalIntensity() string {
	return csi.SelectGraphicRendition(22)
}

// Neither italic, nor black letter.
func NotItalic() string {
	return csi.SelectGraphicRendition(23)
}

// Neither singly nor doubly underlined.
func NotUnderline() string {
	return csi.SelectGraphicRendition(24)
}

// Turn blinking off.
func NotBlinking() string {
	return csi.SelectGraphicRendition(25)
}

// Not crossed out.
func NotCrossedOut() string {
	return csi.SelectGraphicRendition(29)
}

// Disable proportional spacing.
func DisableProportionalSpacing() string {
	return csi.SelectGraphicRendition(50)
}

// Framed: Implemented as "emoji variation selector" in mintty.
func Framed() string {
	return csi.SelectGraphicRendition(51)
}

// Encircled.
func Encircled() string {
	return csi.SelectGraphicRendition(52)
}

// Overlined. Not supported in Terminal.app.
func Overlined() string {
	return csi.SelectGraphicRendition(53)
}

// Neither framed nor encircled.
func NotFramedOrEncircled() string {
	return csi.SelectGraphicRendition(54)
}

// Not overlined.
func NotOverlined() string {
	return csi.SelectGraphicRendition(55)
}
