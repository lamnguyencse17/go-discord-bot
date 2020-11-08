package types

type Response struct {
	OPCODE int `json:"op"`
	DATA map[string]interface{} `json:"d"`
	SEQUENCE int `json:"s"`
	EVENT string `json:"t"`
}

type ConnectionData struct {
	MISC []map[string]interface{} `json:"_trace"`
	INTERVAL int `json:"heartbeat_interval"`
}