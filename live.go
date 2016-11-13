package monobullet

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func wsConnect() {
	u := url.URL{
		Scheme: "wss",
		Host:   streamServer,
		Path:   websocketEndpoint + "/" + config.ApiKey}
	log.Printf("connecting to websocket stream")

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
}
