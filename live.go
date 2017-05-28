package monobullet

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type RealtimeEvent struct {
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
}

var PushChannel = make(chan *Push)

func wsConnect(ctx context.Context) {
	u := url.URL{
		Scheme: "wss",
		Host:   streamServer,
		Path:   websocketEndpoint + "/" + config.APIKey}
	log.Printf("connecting to websocket stream")

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	recv := make(chan *RealtimeEvent)

	go func(c *websocket.Conn, recv chan *RealtimeEvent) {
		for {
			e := new(RealtimeEvent)
			err := c.ReadJSON(e)
			if err != nil {
				log.Fatal(err)
			}
			recv <- e
		}
	}(c, recv)
	lastUpdateTimestamp := int32(time.Now().Unix())

	for {
		select {
		case message := <-recv:
			switch message.Type {
			case "nop":
			case "tickle":
				switch message.Subtype {
				case "push":
					newPushes, err := getPushes(GetPushParams{
						ModifiedAfter: lastUpdateTimestamp,
					})
					if err != nil {
						log.Fatal(err)
					}
					lastUpdateTimestamp = int32(time.Now().Unix())
					for _, push := range newPushes {
						PushChannel <- push
					}
				default:
					fmt.Printf("unhandled tickle subtype: %v\n", message.Subtype)
				}
			case "push":
				// TODO (for ephemerals)
			}
		case <-ctx.Done():
		}
	}

}
