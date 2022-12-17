package services

import (
	"testing"

	"escort-book-escort-profile/types"

	"github.com/stretchr/testify/assert"
)

func TestEmitterServiceAddListener(t *testing.T) {
	emitter := &types.Emitter{
		Name:      "globalEmitter",
		Listeners: nil,
	}
	emitterService := EmitterService{
		Emitter: emitter,
	}

	t.Run("Should add and remove a listener", func(t *testing.T) {
		emitterService.AddListener("test listener", make(chan interface{}))
		assert.Len(t, emitter.Listeners, 1)

		emitterService.RemoveListener("test listener")
		assert.Len(t, emitter.Listeners, 0)
	})
}
