package models

type Message struct {
	user      *User
	Author    string `json:"author"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

func (m *Message) User() *User {
	return m.user
}
