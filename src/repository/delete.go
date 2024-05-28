package repository

import (
	"errors"
	"fmt"
)

func (r *repository) Delete(clientId string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	
	if _, exists := r.store[clientId]; !exists {
		return errors.New(fmt.Sprintf("No message exists for client %s", clientId))
    }
	
	delete(r.store, clientId)
	return nil
}
