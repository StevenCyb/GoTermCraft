package osc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetWindowTitle(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\033]0;abc\007", SetWindowTitle("abc"))
}
