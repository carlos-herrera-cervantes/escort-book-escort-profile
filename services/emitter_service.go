package services

import (
	"escort-book-escort-profile/types"
	"sync"
)

type EmitterService struct{}

var lock = &sync.Mutex{}
var emitter *types.Emitter

func getEmitter() *types.Emitter {
	if emitter == nil {
		lock.Lock()
		defer lock.Unlock()

		if emitter == nil {
			emitter = &types.Emitter{Name: "globalEmitter", Listeners: nil}
		}
	}

	return emitter
}

func (e *EmitterService) AddListener(name string, channel chan interface{}) {
	emitter := getEmitter()

	if emitter.Listeners == nil {
		emitter.Listeners = make(map[string]chan interface{})
	}

	emitter.Listeners[name] = channel
}

func (e *EmitterService) RemoveListener(name string, channel chan interface{}) {
	emitter := getEmitter()

	if listener := emitter.Listeners[name]; listener != nil {
		delete(emitter.Listeners, name)
	}
}

func (e *EmitterService) Emit(name string, message interface{}) {
	emitter := getEmitter()

	if listener := emitter.Listeners[name]; listener != nil {
		go func(listener chan interface{}) {
			listener <- message
		}(listener)
	}
}
