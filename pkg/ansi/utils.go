package ansi

import (
	"fmt"
	"strconv"
	"strings"
)

// Join performs `strings.Join` on the provided array of strings.
func Join(arr ...string) string {
	return strings.Join(arr, "")
}

// Convert byte to string
func B2S(b byte) string {
	if b <= 9 {
		return string(b + '0')
	}

	return fmt.Sprint(b)
}

// Convert HEX string to byte
func HexS2B(s string) byte {
	b, _ := strconv.ParseUint(s, 16, 8)

	return byte(b)
}
