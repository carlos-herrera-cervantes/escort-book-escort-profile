package services

type IEmitterService interface {
	AddListener(name string, channel chan interface{})
	RemoveListener(name string, channel chan interface{})
	Emit(name string, message interface{})
}
