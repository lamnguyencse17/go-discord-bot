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
	interval := client.Session.HearbeatInterval()
	tick := time.NewTicker(time.Second*time.Duration(interval))
	var counter = 0
	for t := range tick.C{
			log.Println("Tick at", t)
			if client.Session.HeartbeatACK() {
				stoppedBeating <- client.Session.HeartbeatACK()
				close(stoppedBeating)
			} else {
				heartbeat(counter)
				counter = counter + 1
			}
	}
}

func heartbeat(counter int) {
	var request HeartbeatRequest
	request.OPCODE = 1
	log.Printf("%v", counter)
	if counter==0 {
		request.DATA = nil
	} else {
		request.DATA = client.Session.Sequence()
	}
	client.Session.Connection().WriteJSON(&request)
	client.Session.ToggleHeartbeatACK()
	log.Printf("done %v heartbeat\n", counter+1)
}
