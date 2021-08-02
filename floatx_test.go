package floatx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInDelta(t *testing.T) {
	assert := assert.New(t)

	assert.True(InDelta(1.23, 1.23, 1e-3))
	assert.True(InDelta(1.23, 1.234, 1e-2))
	assert.False(InDelta(1.23, 1.234, 1e-3))
}

func TestCeilFloat(t *testing.T) {
	assert := assert.New(t)

	assert.InDelta(123.4, CeilFloat(123.4, 1), 1e-2)
	assert.InDelta(123.5, CeilFloat(123.41, 1), 1e-2)
	assert.InDelta(123.5, CeilFloat(123.49, 1), 1e-2)
	assert.InDelta(124, CeilFloat(123.1, 0), 1e-1)
	assert.InDelta(124, CeilFloat(123.9, 0), 1e-1)
	assert.InDelta(124, CeilFloat(123.1, -1), 1e-1)
	assert.InDelta(124, CeilFloat(123.9, -1), 1e-1)
}
