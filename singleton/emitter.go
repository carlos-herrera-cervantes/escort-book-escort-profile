package singleton

import (
	"sync"

	"escort-book-escort-profile/types"
)

var emitter *types.Emitter
var singleEmitter sync.Once

func initEmitter() {
	emitter = &types.Emitter{Name: "globalEmitter", Listeners: nil}
}

func NewEmitter() *types.Emitter {
	singleEmitter.Do(initEmitter)
	return emitter
}
