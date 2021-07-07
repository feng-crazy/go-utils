package ws

import (
	"fmt"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/feng-crazy/go-utils/clock"
)

func TestWsClient(t *testing.T) {
	timeout := clock.GetNowInMilli() + 30000
	deadline := clock.TimeFromUnixMilli(timeout)

	client := &Client{
		Scheme:        "ws",
		Host:          "localhost:9999",
		Path:          "/echo",
		ReadDeadline:  deadline,
		WriteDeadline: deadline,
		conn:          nil,
	}

	err := client.Connect()
	if err != nil {
		logrus.Error(err)
	}
	defer client.DisConnect()

	sendMsg := time.Now().String()
	fmt.Println(sendMsg)
	msgType, msg, err := client.Request(websocket.TextMessage, []byte(sendMsg))
	if err != nil {
		logrus.Error(err)
	}

	fmt.Println(msgType)
	fmt.Println(string(msg))
}
