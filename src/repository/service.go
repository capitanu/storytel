package repository

import "sync"

type Message struct {
    ClientID string `json:"client_id"`
    Content  string `json:"content"`
}

type Repository interface {
	Write(string, string) error
	Update(string, string) error
	Delete(string) error
	Read(string) (string, error)
	List() []Message
}

type repository struct{
	store map[string]Message
    lock sync.Mutex
}

func New() *repository {
	return &repository{
		store: make(map[string]Message),
		lock: sync.Mutex{},
	}
}
