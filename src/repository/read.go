package repository

import (
	"errors"
	"fmt"
)

func (r *repository) Read(clientId string) (string, error) {	
	if message, exists := r.store[clientId]; !exists {
		return "", errors.New(fmt.Sprintf("No message exists for client %s", clientId))
    } else {
		return message.Content, nil
	}
}
