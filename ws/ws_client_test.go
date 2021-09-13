package ws

import (
	"fmt"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func TestWsClient(t *testing.T) {

	client := &Client{
		Scheme: "ws",
		Host:   "localhost:9999",
		Path:   "/echo",

		conn: nil,
	}

	err := client.Connect()
	if err != nil {
		logrus.Error(err)
	}
	defer client.DisConnect()

	for {
		sendMsg := time.Now().String()
		fmt.Println("sendMsg:", sendMsg)
		msgType, msg, err := client.Request(websocket.BinaryMessage, []byte(sendMsg))
		if err != nil {
			logrus.Error(err)
		}

		fmt.Println("msgType:", msgType)
		fmt.Println("recvMsg:", string(msg))

		time.Sleep(3 * time.Second)
		fmt.Println()
	}

}
