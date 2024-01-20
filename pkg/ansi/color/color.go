package color

import "ChromaTerm/pkg/ansi/csi"

// Default foreground color.
func DefaultForegroundColor() string {
	return csi.SelectGraphicRendition(39)
}

// Default background color.
func DefaultBackgroundColor() string {
	return csi.SelectGraphicRendition(49)
}
