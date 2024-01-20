package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	t.Parallel()

	t.Run("ForegroundColor", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\x1b[39m", DefaultForegroundColor())
	})

	t.Run("BackgroundColor", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "\x1b[49m", DefaultBackgroundColor())
	})
}
