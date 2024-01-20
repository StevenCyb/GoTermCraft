package csi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCursor(t *testing.T) {
	t.Parallel()

	t.Run("CursorUp", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[A", CursorUp(1))
		assert.Equal(t, "\033[2A", CursorUp(2))
	})

	t.Run("CursorDown", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[B", CursorDown(1))
		assert.Equal(t, "\033[2B", CursorDown(2))
	})

	t.Run("CursorForward", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[C", CursorForward(1))
		assert.Equal(t, "\033[2C", CursorForward(2))
	})

	t.Run("CursorBack", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[D", CursorBack(1))
		assert.Equal(t, "\033[2D", CursorBack(2))
	})

	t.Run("CursorNextLine", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[E", CursorNextLine(1))
		assert.Equal(t, "\033[2E", CursorNextLine(2))
	})

	t.Run("CursorPreviousLine", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[F", CursorPreviousLine(1))
		assert.Equal(t, "\033[2F", CursorPreviousLine(2))
	})

	t.Run("CursorHorizontalAbsolute", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[G", CursorHorizontalAbsolute(1))
		assert.Equal(t, "\033[2G", CursorHorizontalAbsolute(2))
	})

	t.Run("CursorPosition", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[1;1H", CursorPosition(0, 0))
		assert.Equal(t, "\033[1;1H", CursorPosition(1, 1))
		assert.Equal(t, "\033[2;2H", CursorPosition(2, 2))
	})

	t.Run("CursorSavePosition", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[?25h", ShowCursor())
	})

	t.Run("CursorRestorePosition", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[?25l", HideCursor())
	})
}

func TestErase(t *testing.T) {
	t.Parallel()

	t.Run("EraseScreen", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[0J", EraseInScreen(EraseScreenToEnd))
		assert.Equal(t, "\033[1J", EraseInScreen(EraseScreenToBeginning))
		assert.Equal(t, "\033[2J", EraseInScreen(EraseScreenEntire))
		assert.Equal(t, "\033[3J", EraseInScreen(EraseScreenAndScrollBackBuffer))
		assert.Equal(t, "\033[3J", EraseInScreen(4))
	})

	t.Run("EraseLine", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[0K", EraseInLine(EraseLineToEnd))
		assert.Equal(t, "\033[1K", EraseInLine(EraseLineToBeginning))
		assert.Equal(t, "\033[2K", EraseInLine(EraseLineEntire))
		assert.Equal(t, "\033[2K", EraseInLine(3))
	})
}

func TestScroll(t *testing.T) {
	t.Parallel()

	t.Run("ScrollUp", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[S", ScrollUp(1))
		assert.Equal(t, "\033[2S", ScrollUp(2))
	})

	t.Run("ScrollDown", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "\033[T", ScrollDown(1))
		assert.Equal(t, "\033[2T", ScrollDown(2))
	})
}

func TestSelectGraphicRendition(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\033[0m", SelectGraphicRendition(0))
}
