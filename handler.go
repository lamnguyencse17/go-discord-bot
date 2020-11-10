package main

import (
	"encoding/json"
	"github.com/lamnguyencse17/go-discord-bot/handlers"

	//"github.com/lamnguyencse17/go-discord-bot/handlers"
	client "github.com/lamnguyencse17/go-discord-bot/session"
	types "github.com/lamnguyencse17/go-discord-bot/types"
	"log"
)

func MessageListener(done chan bool, startBeating chan bool){
	startBeating <- false
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
		processOPCODE(response, response.OPCODE, startBeating)
	}
}

func processOPCODE (response types.Response, OPCODE int, startBeating chan bool){
	switch OPCODE {
	case 10:
		client.Session.InitHeartbeatAck()
		client.Session.SetHeartbeatInterval(response.DATA.INTERVAL)
		handlers.InitConnection(response)
		startBeating <- true
		close(startBeating)
	case 11:
		log.Printf("Acknowledged Heartbeat")
		client.Session.ToggleHeartbeatACK()
		log.Printf("waiting?: %v", client.Session.HeartbeatACK())
	//case 0:
	//	processEVENT(response)
	}
}

func processEVENT (response types.Response){
	switch response.EVENT {
	case "READY":
		handlers.ReadyConnection(response)
	case "GUILD_CREATE":
	}
}