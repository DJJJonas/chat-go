package models

import (
	"log"
	"net/http"
	"strings"

	"github.com/DJJJonas/chat-go/utils"
)

type ChatRoom struct {
	Users       map[uint32]*User
	UserCount   uint
	ChatChannel *ChatChannel
}

func (cr *ChatRoom) ServerHandler(w http.ResponseWriter, r *http.Request) {
	userConnection, err := utils.UpgradeConnection(w, r)
	if err != nil {
		utils.WEUpgradeRequired(w, err)
		return
	}

	nickname := strings.TrimSpace(r.URL.Query().Get("nickname"))
	u := &User{
		ID:          uint32(cr.UserCount),
		Name:        nickname,
		chatChannel: cr.ChatChannel,
		conn:        userConnection,
	}

	cr.ChatChannel.JoinChannel <- u // Emit the join channel event

	u.Mainloop()
}

// This event handler will handle user-sent events such as
// "join chat", "leave chat" and "send a message" event
func (cr *ChatRoom) HandleUserEvents() {
	for {
		select {
		case u := <-cr.ChatChannel.JoinChannel: // User is joining the chat
			cr.UserCount++
			// The user's id will be the number of connection the server
			// had so far
			cr.Users[u.ID] = u
			log.Printf("User %s has connected with id: %d", u.Name, u.ID)

		case u := <-cr.ChatChannel.LeaveChannel: // User is leaving the chat
			delete(cr.Users, u.ID)
			log.Printf("User %s has connected with id: %d", u.Name, u.ID)

		case m := <-cr.ChatChannel.MessageChannel: // User sent a message
			// Broadcast the message to all users...
			for _, v := range cr.Users {
				// ...except for the one that sent the message
				if v.ID != m.User().ID {
					v.WriteMessage(m)
				}
			}
		}
	}
}
