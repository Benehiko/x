package stringsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultIfEmpty(t *testing.T) {
	assert.Equal(t, DefaultIfEmpty("", "default"), "default")
	assert.Equal(t, DefaultIfEmpty("custom", "default"), "custom")
}
