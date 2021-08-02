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
