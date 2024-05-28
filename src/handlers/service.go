package handlers

type Handler interface {
	Init()
}

type handler struct{}

func New() *handler {
	return &handler{}
}
