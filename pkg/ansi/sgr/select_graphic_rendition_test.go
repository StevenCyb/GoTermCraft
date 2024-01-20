package sgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[0m", Reset())
}

func TestBold(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[1m", Bold())
}

func TestFaint(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[2m", Faint())
}

func TestItalic(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[3m", Italic())
}

func TestUnderline(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[4m", Underline())
}

func TestBlinkSlow(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[5m", BlinkSlow())
}

func TestBlinkRapid(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[6m", BlinkRapid())
}

func TestImageNegative(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[7m", ImageNegative())
}

func TestStrike(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[9m", Strike())
}

func TestDoubleUnderline(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[21m", DoubleUnderline())
}

func TestNormalIntensity(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[22m", NormalIntensity())
}

func TestNotItalic(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[23m", NotItalic())
}

func TestNotUnderline(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[24m", NotUnderline())
}

func TestNotBlinking(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[25m", NotBlinking())
}

func TestNotCrossedOut(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[29m", NotCrossedOut())
}

func TestDisableProportionalSpacing(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[50m", DisableProportionalSpacing())
}

func TestFramed(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[51m", Framed())
}

func TestEncircled(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[52m", Encircled())
}

func TestOverlined(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[53m", Overlined())
}

func TestNotFramedOrEncircled(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[54m", NotFramedOrEncircled())
}

func TestNotOverlined(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "\x1b[55m", NotOverlined())
}

