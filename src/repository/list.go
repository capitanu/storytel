package repository

func (r *repository) List() []Message {
	var messages []Message
	for _, message := range r.store {
		messages = append(messages, message)
	}
	return messages
}
