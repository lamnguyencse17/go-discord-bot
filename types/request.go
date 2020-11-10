package types

type IdentifyRequest struct {
	OPCODE int `json:"op"`
	DATA IdentifyData `json:"d"`
}

type IdentifyData struct {
	TOKEN string `json:"token"`
	INTENTS int `json:"intents"`
	PROPERTIES IdentifyDataProperties `json:"properties"`
	//COMPRESS bool `json:"compress,omitempty"`
	//THRESHOLD int `json:"large_threshold,omitempty"`
}

type IdentifyDataProperties struct {
	OS string `json:"$os"`
	BROWSER string `json:"$browser"`
	DEVICE string `json:"$device"`
}