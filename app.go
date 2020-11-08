package main

import (
	"flag"
	"github.com/gorilla/websocket"
	client "github.com/lamnguyencse17/go-discord-bot/session"
	"log"
	"net/url"
	"os"
	"os/signal"
)

var addr = flag.String("addr", "gateway.discord.gg", "Discord Gateway API")

func main() {
	client.Session.SetGateway(*addr)
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/"}
	q := u.Query()
	q.Set("v", "8")
	q.Set("encoding", "json")
	u.RawQuery = q.Encode()
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	client.Session.SetConnection(c)
	if err != nil {
		log.Fatal("dial: ", err)
	}
	defer c.Close()
	done := make(chan bool)
	go EventHandler(done)
	for {
		select {
		case closed := <-done:
			if !closed {
				log.Printf("Done")
				return
			}
		}
	}
}