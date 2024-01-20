package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test4Bit(t *testing.T) {
	t.Parallel()

	t.Run("Foreground", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[30m", Foreground4Bit(0))
		assert.Equal(t, "\033[30m", Foreground4Bit(ForegroundBlack))
		assert.Equal(t, "\033[37m", Foreground4Bit(ForegroundWhite))
		assert.Equal(t, "\033[30m", Foreground4Bit(38))
		assert.Equal(t, "\033[90m", Foreground4Bit(ForegroundBrightBlack))
		assert.Equal(t, "\033[97m", Foreground4Bit(ForegroundBrightWhite))
		assert.Equal(t, "\033[30m", Foreground4Bit(98))
	})

	t.Run("Background", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[40m", Background4Bit(0))
		assert.Equal(t, "\033[40m", Background4Bit(BackgroundBlack))
		assert.Equal(t, "\033[47m", Background4Bit(BackgroundWhite))
		assert.Equal(t, "\033[40m", Background4Bit(48))
		assert.Equal(t, "\033[100m", Background4Bit(BackgroundBrightBlack))
		assert.Equal(t, "\033[107m", Background4Bit(BackgroundBrightWhite))
		assert.Equal(t, "\033[40m", Background4Bit(108))
	})

	t.Run("ForegroundAndBackground", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[30;40m", Color4Bit(0, 0))
		assert.Equal(t, "\033[30;40m", Color4Bit(ForegroundBlack, BackgroundBlack))
		assert.Equal(t, "\033[37;47m", Color4Bit(ForegroundWhite, BackgroundWhite))
		assert.Equal(t, "\033[30;40m", Color4Bit(38, 48))
		assert.Equal(t, "\033[90;100m", Color4Bit(ForegroundBrightBlack, BackgroundBrightBlack))
		assert.Equal(t, "\033[97;107m", Color4Bit(ForegroundBrightWhite, BackgroundBrightWhite))
		assert.Equal(t, "\033[30;40m", Color4Bit(98, 108))
	})
}
