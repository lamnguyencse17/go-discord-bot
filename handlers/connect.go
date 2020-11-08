package handlers

import (
	client "github.com/lamnguyencse17/go-discord-bot/session"
	types "github.com/lamnguyencse17/go-discord-bot/types"
	"github.com/mitchellh/mapstructure"
	"log"
)

func ConnectionHandler (response types.Response){
	var connectionData types.ConnectionData
	if err:= mapstructure.Decode(response, &connectionData); err != nil{
		log.Fatalf("decode map: %s", err)
	}
	client.Session.SetHeartbeatInterval(connectionData.INTERVAL)
	client.Session.SetSequence(response.SEQUENCE)
	log.Printf("Session: %v", client.Session)
}
