package repository

import (
	"errors"
	"fmt"
)

func (r *repository) Update(clientId string, content string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, exists := r.store[clientId]; !exists {
		// Ehm, potentially a security problem for other systems, but should be fine here
        return errors.New(fmt.Sprintf("No message exists for client %s", clientId))
    }

	message := Message{
		ClientID: clientId,
		Content: content,
	}
	
	r.store[clientId] = message
	return nil
}
