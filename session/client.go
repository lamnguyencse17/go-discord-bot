package session

import (
	"github.com/gorilla/websocket")

type Client struct {
	gateway           string
	sequence          int
	token             string
	heartbeatInterval int
	connection *websocket.Conn
}

var Session Client

func (client *Client) SetGateway(gateway string) {
	client.gateway = gateway
}

func (client *Client) SetConnection(connection *websocket.Conn){
	client.connection = connection
}

func (client *Client) SetSequence(sequence int) {
	client.sequence = sequence
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