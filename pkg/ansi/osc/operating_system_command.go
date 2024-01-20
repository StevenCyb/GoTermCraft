package osc

// SetWindowTitle returns the ANSI escape code
// for setting the window title.
func SetWindowTitle(title string) string {
	return "\033]0;" + title + "\007"
}
