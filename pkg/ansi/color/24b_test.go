package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test24Bit(t *testing.T) {
	t.Parallel()

	t.Run("Foreground", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[38;2;0;0;0m", Foreground24Bit(0, 0, 0))
	})

	t.Run("Background", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[48;2;0;0;0m", Background24Bit(0, 0, 0))
	})

	t.Run("ForegroundHex", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[38;2;0;0;0m", Foreground24BitHex("000000"))
	})

	t.Run("BackgroundHex", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[48;2;0;0;0m", Background24BitHex("000000"))
	})
}

func TestHexConvert(t *testing.T) {
	t.Parallel()

	r, g, b := hex2rgb("#000000")
	assert.Equal(t, []byte{0, 0, 0}, []byte{r, g, b})

	r, g, b = hex2rgb("#000")
	assert.Equal(t, []byte{0, 0, 0}, []byte{r, g, b})

	r, g, b = hex2rgb("#000000000")
	assert.Equal(t, []byte{0, 0, 0}, []byte{r, g, b})

	r, g, b = hex2rgb("#ff0000")
	assert.Equal(t, []byte{255, 0, 0}, []byte{r, g, b})

	r, g, b = hex2rgb("#00FF00")
	assert.Equal(t, []byte{0, 255, 0}, []byte{r, g, b})

	r, g, b = hex2rgb("#0000FF")
	assert.Equal(t, []byte{0, 0, 255}, []byte{r, g, b})
}
