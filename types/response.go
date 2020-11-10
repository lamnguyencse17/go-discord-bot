package types

type Response struct {
	SEQUENCE int `json:"s"`
	EVENT string `json:"t"`
	OPCODE int `json:"op"`
	DATA struct {
		INTERVAL int `json:"heartbeat_interval"`
		VERSION int `json:"v"`
		USER User `json:"user"`
	} `json:"d"`
}