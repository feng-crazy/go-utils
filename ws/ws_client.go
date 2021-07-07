package ws

import (
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Client struct {
	Scheme        string
	Host          string // host or host:port
	Path          string // path (relative paths may omit leading slash)
	ReadDeadline  time.Time
	WriteDeadline time.Time
	conn          *websocket.Conn
}

func NewClient(scheme, host, path string) Client {
	return Client{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
}

func (t *Client) Connect() error {
	var err error

	u := url.URL{Scheme: t.Scheme, Host: t.Host, Path: t.Path}
	t.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logrus.Error("dial err:", err)
		return err
	}

	err = t.conn.SetReadDeadline(t.ReadDeadline)
	if err != nil {
		logrus.Error("SetReadDeadline err:", err)
		return err
	}

	err = t.conn.SetWriteDeadline(t.WriteDeadline)
	if err != nil {
		logrus.Error("SetWriteDeadline err:", err)
		return err
	}
	return nil
}

func (t *Client) Request(reqMsgType int, reqMsg []byte) (msgType int, msg []byte, err error) {
	err = t.conn.WriteMessage(reqMsgType, reqMsg)
	if err != nil {
		logrus.Error("WriteMessage err:", err)
		return
	}

	msgType, msg, err = t.conn.ReadMessage()
	if err != nil {
		logrus.Error("ReadMessage err:", err)
		return
	}
	return
}

func (t *Client) RecvByte() (messageType int, message []byte, err error) {
	messageType, message, err = t.conn.ReadMessage()
	if err != nil {
		logrus.Error("ReadMessage err:", err)
		return
	}
	return
}

func (t *Client) SendByte(messageType int, message []byte) (err error) {
	err = t.conn.WriteMessage(messageType, []byte(message))
	if err != nil {
		logrus.Error("WriteMessage err:", err)
		return
	}
	return
}

func (t *Client) RecvString() (messageType int, message string, err error) {
	messageType, p, err := t.conn.ReadMessage()
	if err != nil {
		logrus.Error("ReadMessage err:", err)
		return
	}
	message = string(p)
	return
}

func (t *Client) SendString(messageType int, message string) (err error) {
	err = t.conn.WriteMessage(messageType, []byte(message))
	if err != nil {
		logrus.Error("WriteMessage err:", err)
		return
	}
	return
}

func (t *Client) DisConnect() {
	err := t.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logrus.Println("write close:", err)
		return
	}
	_ = t.conn.Close()
}
