package singleton

import (
	"testing"

	"escort-book-escort-profile/types"

	"github.com/stretchr/testify/assert"
)

func TestNewEmitter(t *testing.T) {
	t.Run("Should return a pointer to emitter", func(t *testing.T) {
		emitter := NewEmitter()
		assert.IsType(t, &types.Emitter{}, emitter)
	})
}
