# ANSI
The `ansi` package provides functions to generate ANSI escape codes, which are used to control the formatting, color, and other output options on text terminals.

## Sub-packages

### Color
The color sub-package provides functions to set the foreground and background colors of the text. 
It supports both 4-bit and 8-bit color modes, as well as 24-bit true color.
* 4-bit color: Refer to 4b.md for more details.
* 8-bit color: Refer to 8b.md for more details.
* 24-bit color: Use the `Foreground24` and `Background24` functions to set the foreground and background colors, respectively. Each function takes three `byte` parameters representing the red, green, and blue color values.

For the color palettes check out:
* [4-bit](color/4b.md)
* [8-bit](color/8b.md)
* 24-bit uses RGB

### CSI (Control Sequence Introducer)
The `csi` sub-package provides functions to control the cursor position and visibility, and to erase parts of the screen. Here are some of the functions:
* `CursorUp`, `CursorDown`, `CursorForward`, `CursorBack`: Move the cursor a specified number of cells in the given direction.
* `CursorNextLine`, `CursorPreviousLine`: Move the cursor to the beginning of the line a specified number of lines down or up.
* `CursorPosition`: Move the cursor to a specific row and column.
* `ShowCursor`, `HideCursor`: Show or hide the cursor.
* `EraseScreen`: Erase parts of the screen. It takes a parameter of type `EraseScreen`, which can be one of the following constants: `EraseScreenToEnd`, `EraseScreenToBeginning`, `EraseScreenEntire`, `EraseScreenAndScrollBackBuffer`.

### OSC (Operating System Command)
The `osc` sub-package provides functions to send operating system commands.

### SGR (Select Graphic Rendition)
The `sgr` sub-package provides functions to set graphic rendition attributes, such as bold, underline, blink, and reverse video.
