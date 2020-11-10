package handlers

import (
	client "github.com/lamnguyencse17/go-discord-bot/session"
	types "github.com/lamnguyencse17/go-discord-bot/types"
	"log"
	"os"
)

func InitConnection(response types.Response){
	//var connectionData = types.ConnectionData{}
	//	////data := response.DATA
	//	//if err:= mapstructure.Decode(response, &connectionData); err != nil{
	//	//	log.Fatalf("decode map: %s", err)
	//	//}
	//	//client.Session.SetHeartbeatInterval(connectionData.INTERVAL)
	//var initResponse = response.ParseInitResponse()
	log.Printf("init Response: %v", response)
	client.Session.SetSequence(response.SEQUENCE)
	IdentifyConnection()
}

func IdentifyConnection(){
	var IdentifyRequest types.IdentifyRequest
	var IdentifyData types.IdentifyData
	var IdentifyDataProperties types.IdentifyDataProperties
	IdentifyRequest.OPCODE = 2
	IdentifyData.TOKEN = os.Getenv("discord_token")
	IdentifyData.INTENTS = 513
	IdentifyDataProperties.BROWSER = "zodiac3011"
	IdentifyDataProperties.DEVICE = "zodiac3011"
	IdentifyDataProperties.OS = "browser"
	IdentifyData.PROPERTIES = IdentifyDataProperties
	IdentifyRequest.DATA = IdentifyData
	client.Session.Connection().WriteJSON(&IdentifyRequest)
}

func ReadyConnection(response types.Response){
	//var ReadyResponse types.ReadyResponse
	//ReadyResponse = response.ParseReadyResponse()
	client.Session.SetSequence(response.SEQUENCE)
	client.Session.SetSessionID(response.DATA.USER.Session_id)
	client.Session.SetUser(response.DATA.USER)
}