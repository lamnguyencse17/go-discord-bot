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
	interval := client.Session.HearbeatInterval()
	tick := time.Tick(time.Millisecond*time.Duration(interval))
	counter := 0
	var request HeartbeatRequest
	heartbeat(request, counter)
	for {
		select {
		case <-tick:
			heartbeat(request, counter)
		}
	}
}

func heartbeat(request HeartbeatRequest, counter int) {
	log.Printf("Beating")
	request.OPCODE = 1
	if counter==0 {
		request.DATA = nil
	} else {
		*request.DATA = client.Session.Sequence()
	}
	client.Session.Connection().WriteJSON(&request)
	counter = counter + 1
	log.Printf("done %v heartbeat\n", counter)
}
