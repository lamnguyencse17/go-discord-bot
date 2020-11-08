package handlers

import (
	client "github.com/lamnguyencse17/go-discord-bot/session"
	types "github.com/lamnguyencse17/go-discord-bot/types"
	"github.com/mitchellh/mapstructure"
	"log"
)

func InitConnection(response types.Response){
	var connectionData types.ConnectionData
	if err:= mapstructure.Decode(response, &connectionData); err != nil{
		log.Fatalf("decode map: %s", err)
	}
	client.Session.SetHeartbeatInterval(connectionData.INTERVAL)
	client.Session.SetSequence(response.SEQUENCE)
	IdentifyConnection()
}

func IdentifyConnection(){
	var IdentifyRequest types.IdentifyRequest
	var IdentifyData types.IdentifyData
	var IdentifyDataProperties types.IdentifyDataProperties
	IdentifyRequest.OPCODE = 2
	IdentifyData.TOKEN = "NjIxMjMzMDEyMDkyNDM2NTAw.XXiWVA.t8wyUeY3ToOp9qWT2f98iqkVSDk"
	IdentifyData.INTENTS = 513
	IdentifyDataProperties.BROWSER = "zodiac3011"
	IdentifyDataProperties.DEVICE = "zodiac3011"
	IdentifyDataProperties.OS = "browser"
	IdentifyData.PROPERTIES = IdentifyDataProperties
	IdentifyRequest.DATA = IdentifyData
	client.Session.Connection().WriteJSON(&IdentifyRequest)
}