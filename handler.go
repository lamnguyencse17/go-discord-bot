package main

import (
	"encoding/json"
	"github.com/lamnguyencse17/go-discord-bot/handlers"
	client "github.com/lamnguyencse17/go-discord-bot/session"
	types "github.com/lamnguyencse17/go-discord-bot/types"
	"log"
)

func EventHandler (done chan bool){
	defer close(done)

	var response = types.Response{}
	for {
		_, message, err := client.Session.Connection().ReadMessage()
		if err != nil {
			log.Printf("read: %s", err)
			return
		}

		if err = json.Unmarshal(message, &response); err != nil{
			log.Printf("json: %s", err)
			return
		}
		log.Printf("raw: %s\n", message)
		log.Printf("recv: %v\n", response)
		stoppedBeating := make(chan bool)
		if response.OPCODE == 10{
			handlers.ConnectionHandler(response)
			go Heartbeat(stoppedBeating)
		}
	}
}
