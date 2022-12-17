package services

//go:generate mockgen -destination=./mocks/iemitter_service.go -package=mocks --build_flags=--mod=mod . IEmitterService
type IEmitterService interface {
	AddListener(name string, channel chan interface{})
	RemoveListener(name string)
	Emit(name string, message interface{})
}
