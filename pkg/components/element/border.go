package element

type BorderStyle struct {
	TopLeft     rune
	TopRight    rune
	BottomLeft  rune
	BottomRight rune
	Top         rune
	Bottom      rune
	Left        rune
	Right       rune
}

func DefaultBorderStyle() BorderStyle {
	return BorderStyle{
		TopLeft:     '┌',
		TopRight:    '┐',
		BottomLeft:  '└',
		BottomRight: '┘',
		Top:         '─',
		Bottom:      '─',
		Left:        '│',
		Right:       '│',
	}
}

func HashTagBorderStyle() BorderStyle {
	return BorderStyle{
		TopLeft:     '#',
		TopRight:    '#',
		BottomLeft:  '#',
		BottomRight: '#',
		Top:         '#',
		Bottom:      '#',
		Left:        '#',
		Right:       '#',
	}
}
