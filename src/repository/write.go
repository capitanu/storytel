package repository

import (
	"errors"
)

func (r *repository) Write(clientId string, content string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	
	if _, exists := r.store[clientId]; exists {
        return errors.New("Message already exists")
    }

	message := Message{
		ClientID: clientId,
		Content: content,
	}
	
	r.store[clientId] = message
	return nil
}
