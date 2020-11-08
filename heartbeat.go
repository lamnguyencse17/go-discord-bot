package main

import (
	client "github.com/lamnguyencse17/go-discord-bot/session"
	"log"
	"time"
)

type HeartbeatRequest struct{
	OPCODE int `json:"op"`
	DATA *int `json:"d"`
}

func Heartbeat(stoppedBeating chan bool) {
	defer close(stoppedBeating)
	log.Printf("Stating Heartbeat")
	//interval := client.Session.HearbeatInterval()
	tick := time.NewTicker(time.Second*2)
	var counter = 0
	var request HeartbeatRequest
	for t := range tick.C{
			log.Println("Tick at", t)
			if client.Session.HeartbeatACK() {
				stoppedBeating <- client.Session.HeartbeatACK()
				close(stoppedBeating)
			} else {
				heartbeat(request, counter)
				counter = counter + 1
			}
	}
}

func heartbeat(request HeartbeatRequest, counter int) {
	request.OPCODE = 1
	if counter==0 {
		request.DATA = nil
	} else {
		*request.DATA = client.Session.Sequence()
	}
	client.Session.Connection().WriteJSON(&request)
	client.Session.ToggleHeartbeatACK()
	log.Printf("done %v heartbeat\n", counter+1)
}
