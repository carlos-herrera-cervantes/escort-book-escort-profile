package services

import "escort-book-escort-profile/types"

type EmitterService struct {
	Emitter *types.Emitter
}

func (e *EmitterService) AddListener(name string, channel chan interface{}) {
	if e.Emitter.Listeners == nil {
		e.Emitter.Listeners = make(map[string]chan interface{})
	}

	e.Emitter.Listeners[name] = channel
}

func (e *EmitterService) RemoveListener(name string) {
	if listener := e.Emitter.Listeners[name]; listener != nil {
		delete(e.Emitter.Listeners, name)
	}
}

func (e *EmitterService) Emit(name string, message interface{}) {
	if listener := e.Emitter.Listeners[name]; listener != nil {
		go func(listener chan interface{}) {
			listener <- message
		}(listener)
	}
}
