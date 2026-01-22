package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyAnyValue(t *testing.T) {
	buf := []byte{0x80, 0x84, 0x80, 0x00, 0x0A, 0xC4}
	length, pos, _ := ConsumeLen(buf, 0)
	assert.Equal(t, 516, pos)
	assert.Equal(t, 512, length)
}
