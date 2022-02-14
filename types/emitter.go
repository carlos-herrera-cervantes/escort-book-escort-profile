package types

type Emitter struct {
	Name      string
	Listeners map[string]chan interface{}
}
