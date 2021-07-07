package ws

import (
	"flag"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var addr = flag.String("addr", "localhost:9999", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			logrus.Println("read:", err)
			break
		}
		logrus.Printf("recv: %s, mt:%+v", message, mt)
		err = c.WriteMessage(mt, message)
		if err != nil {
			logrus.Println("write:", err)
			break
		}
	}
}

func TestGorillaServer(t *testing.T) {
	flag.Parse()
	http.HandleFunc("/echo", echo)
	logrus.Fatal(http.ListenAndServe(*addr, nil))
}
