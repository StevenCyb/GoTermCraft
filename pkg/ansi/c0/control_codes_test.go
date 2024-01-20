package c0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBell(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\a", Bell())
}
