package ansi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "abc", Join("a", "b", "c"))
}

func TestB2S(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "0", B2S(0))
	assert.Equal(t, "9", B2S(9))
	assert.Equal(t, "10", B2S(10))
}

func TestS2B(t *testing.T) {
	t.Parallel()

	assert.Equal(t, byte(0), HexS2B(""))
	assert.Equal(t, byte(0), HexS2B("0"))
	assert.Equal(t, byte(9), HexS2B("9"))
	assert.Equal(t, byte(10), HexS2B("A"))
	assert.Equal(t, byte(15), HexS2B("F"))
	assert.Equal(t, byte(255), HexS2B("1000"))
}
