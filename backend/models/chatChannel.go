package models

// ChatChannel will store user-sent events
type ChatChannel struct {
	JoinChannel    chan *User    // User to be added to chat
	LeaveChannel   chan *User    // User ID to be removed from chat
	MessageChannel chan *Message // Message to be broadcasted
}
