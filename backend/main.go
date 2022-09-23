package main

import (
	"log"
	"net/http"

	"github.com/DJJJonas/chat-go/models"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {

	r := &models.ChatRoom{
		Users:     make(map[uint32]*models.User),
		UserCount: 1,
		ChatChannel: &models.ChatChannel{
			JoinChannel:    make(chan *models.User),
			LeaveChannel:   make(chan *models.User),
			MessageChannel: make(chan *models.Message),
		},
	}

	// Endpoint to handle every user connection
	http.HandleFunc("/chat", r.ServerHandler)

	go r.HandleUserEvents()

	log.Printf("Listening on %s:%s 👂", HOST, PORT)
	log.Fatal(http.ListenAndServe(HOST+":"+PORT, nil))
}
