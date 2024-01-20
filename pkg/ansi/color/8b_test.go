package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test8Bit(t *testing.T) {
	t.Parallel()

	t.Run("Foreground", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[38;5;0m", Foreground8Bit(0))
	})

	t.Run("Background", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[48;5;0m", Background8Bit(0))
	})
}
