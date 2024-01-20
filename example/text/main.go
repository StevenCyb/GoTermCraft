package main

import (
	"ChromaTerm/pkg/ansi"
	"ChromaTerm/pkg/ansi/color"
	"ChromaTerm/pkg/components/text"
	"fmt"
)

func main() {
	fmt.Println(ansi.Join(
		text.Text("Black").Foreground4Bit(color.ForegroundBlack).String(),
		text.Text(" Red").Foreground4Bit(color.ForegroundRed).String(),
		text.Text(" Green").Foreground4Bit(color.ForegroundGreen).String(),
		text.Text(" Yellow").Foreground4Bit(color.ForegroundYellow).String(),
		text.Text(" Blue").Foreground4Bit(color.ForegroundBlue).String(),
		text.Text(" Magenta").Foreground4Bit(color.ForegroundMagenta).String(),
		text.Text(" Cyan").Foreground4Bit(color.ForegroundCyan).String(),
		text.Text(" White").Foreground4Bit(color.ForegroundWhite).String(),
	))

	fmt.Println(ansi.Join(
		text.Text("Black").Foreground4Bit(color.ForegroundBrightBlack).String(),
		text.Text(" Red").Foreground4Bit(color.ForegroundBrightRed).String(),
		text.Text(" Green").Foreground4Bit(color.ForegroundBrightGreen).String(),
		text.Text(" Yellow").Foreground4Bit(color.ForegroundBrightYellow).String(),
		text.Text(" Blue").Foreground4Bit(color.ForegroundBrightBlue).String(),
		text.Text(" Magenta").Foreground4Bit(color.ForegroundBrightMagenta).String(),
		text.Text(" Cyan").Foreground4Bit(color.ForegroundBrightCyan).String(),
		text.Text(" White").Foreground4Bit(color.ForegroundBrightWhite).String(),
	))

	fmt.Println(ansi.Join(
		text.Text("Black").Background4Bit(color.BackgroundBlack).StringWithReset(), " ",
		text.Text("Red").Background4Bit(color.BackgroundRed).StringWithReset(), " ",
		text.Text("Green").Background4Bit(color.BackgroundGreen).StringWithReset(), " ",
		text.Text("Yellow").Background4Bit(color.BackgroundYellow).StringWithReset(), " ",
		text.Text("Blue").Background4Bit(color.BackgroundBlue).StringWithReset(), " ",
		text.Text("Magenta").Background4Bit(color.BackgroundMagenta).StringWithReset(), " ",
		text.Text("Cyan").Background4Bit(color.BackgroundCyan).StringWithReset(), " ",
		text.Text("White").Background4Bit(color.BackgroundWhite).StringWithReset(),
	))

	fmt.Println(ansi.Join(
		text.Text("Black").Background4Bit(color.BackgroundBrightBlack).StringWithReset(), " ",
		text.Text("Red").Background4Bit(color.BackgroundBrightRed).StringWithReset(), " ",
		text.Text("Green").Background4Bit(color.BackgroundBrightGreen).StringWithReset(), " ",
		text.Text("Yellow").Background4Bit(color.BackgroundBrightYellow).StringWithReset(), " ",
		text.Text("Blue").Background4Bit(color.BackgroundBrightBlue).StringWithReset(), " ",
		text.Text("Magenta").Background4Bit(color.BackgroundBrightMagenta).StringWithReset(), " ",
		text.Text("Cyan").Background4Bit(color.BackgroundBrightCyan).StringWithReset(), " ",
		text.Text("White").Background4Bit(color.BackgroundBrightWhite).StringWithReset(),
	))

	fmt.Println(ansi.Join(
		text.Text("Black").Background24Bit(0, 0, 0).StringWithReset(), " ",
		text.Text("Red").Background24Bit(255, 0, 0).StringWithReset(), " ",
		text.Text("Green").Background24Bit(0, 255, 0).StringWithReset(), " ",
		text.Text("Yellow").Background24Bit(255, 255, 0).StringWithReset(), " ",
		text.Text("Blue").Background24Bit(0, 0, 255).StringWithReset(), " ",
		text.Text("Magenta").Background24Bit(255, 0, 255).StringWithReset(), " ",
		text.Text("Cyan").Background24Bit(0, 255, 255).StringWithReset(), " ",
		text.Text("White").Background24Bit(255, 255, 255).StringWithReset(),
	))

	fmt.Println(ansi.Join(
		text.Text("██████").Foreground24BitHex("#FFD6A5").StringWithReset(),
		text.Text("███████").Foreground24BitHex("#FFADAD").StringWithReset(),
		text.Text("███████").Foreground24BitHex("#CAFFBF").StringWithReset(),
		text.Text("██████").Foreground24BitHex("#FDFFB6").StringWithReset(),
		text.Text("███████").Foreground24BitHex("#2C2C54").StringWithReset(),
		text.Text("███████").Foreground24BitHex("#474787").StringWithReset(),
		text.Text("██████").Foreground24BitHex("#AAABB8").StringWithReset(),
	))

	fmt.Println(ansi.Join(
		text.Text("Bold").Bold().StringWithReset(), " ",
		text.Text("Faint").Faint().StringWithReset(), " ",
		text.Text("Italic").Italic().StringWithReset(), " ",
		text.Text("Underline").Underline().StringWithReset(), " ",
		text.Text("Strike").Strike().StringWithReset(), " ",
		text.Text("D.Underline").DoubleUnderline().StringWithReset(), " ",
	))
}
