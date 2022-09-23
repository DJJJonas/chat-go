package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	chatChannel *ChatChannel
	conn        *websocket.Conn
}

func (u *User) Mainloop() {
	for {
		if _, messageInBytes, err := u.conn.ReadMessage(); err != nil {
			log.Printf("Disconnecting %s", u.Name)
			break
		} else {
			message := &Message{}
			if err := json.Unmarshal(messageInBytes, message); err != nil {
				log.Printf("JSON error: %s", err.Error())
			} else {
				message.user = u
				message.CreatedAt = time.Now().Format("15:04 PM")
				u.chatChannel.MessageChannel <- message // Sent event of message
			}
		}
	}

	u.chatChannel.LeaveChannel <- u
}

func (u *User) WriteMessage(message *Message) {
	message.Author = message.user.Name
	if bytes, err := json.Marshal(message); err != nil {
		log.Printf("JSON Error: %s", err.Error())
	} else {
		u.conn.WriteMessage(websocket.TextMessage, bytes)
	}
}
