package session

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	mu sync.Mutex
	gateway           string
	sequence          int
	token             string
	heartbeatInterval int
	connection *websocket.Conn
	waitingHeartbeatACK bool
}

var Session Client

func (client *Client) SetGateway(gateway string) {
	client.gateway = gateway
}

func (client *Client) SetConnection(connection *websocket.Conn){
	client.connection = connection
}

func (client *Client) SetSequence(sequence int) {
	client.mu.Lock()
	client.sequence = sequence
	client.mu.Unlock()
}

func (client *Client) InitHeartbeatAck() {
	client.mu.Lock()
	client.waitingHeartbeatACK = false
	client.mu.Unlock()
}

func (client *Client) ToggleHeartbeatACK() {
	client.mu.Lock()
	client.waitingHeartbeatACK = !client.waitingHeartbeatACK
	client.mu.Unlock()
}

func (client *Client) SetToken(token string) {
	client.token = token
}

func (client *Client) SetHeartbeatInterval(interval int) {
	client.heartbeatInterval = interval
}

func (client *Client) Gateway() string {
	return client.gateway
}

func (client *Client) Sequence() int {
	client.mu.Lock()
	defer client.mu.Unlock()
	return client.sequence
}

func (client *Client) Token() string {
	return client.token
}

func (client *Client) Connection() *websocket.Conn {
	return client.connection
}

func (client *Client) HearbeatInterval() int {
	return client.heartbeatInterval
}

func (client *Client) HeartbeatACK() bool{
	client.mu.Lock()
	defer client.mu.Unlock()
	return client.waitingHeartbeatACK
}