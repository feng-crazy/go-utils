package ws

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ConnectionOptions struct {
	UseCompression bool
	UseSSL         bool
	Proxy          func(*http.Request) (*url.URL, error)
	Subprotocols   []string
}

type Client struct {
	Scheme string
	Host   string // host or host:port
	Path   string // path (relative paths may omit leading slash)
	Conn   *websocket.Conn

	WebsocketDialer   *websocket.Dialer
	ConnectionOptions ConnectionOptions
	RequestHeader     http.Header
}

func NewClient(scheme, host, path string) *Client {
	return &Client{
		Scheme: scheme,
		Host:   host,
		Path:   path,
		Conn:   nil,

		RequestHeader: http.Header{},
		ConnectionOptions: ConnectionOptions{
			UseCompression: false,
			UseSSL:         true,
		},
		WebsocketDialer: &websocket.Dialer{},
	}
}

func (t *Client) setConnectionOptions() {
	t.WebsocketDialer.EnableCompression = t.ConnectionOptions.UseCompression
	t.WebsocketDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: t.ConnectionOptions.UseSSL}
	t.WebsocketDialer.Proxy = t.ConnectionOptions.Proxy
	t.WebsocketDialer.Subprotocols = t.ConnectionOptions.Subprotocols
}

func (t *Client) Connect() error {
	var err error

	// u := url.URL{Scheme: t.Scheme, Host: t.Host, Path: t.Path}
	addr := fmt.Sprintf("%s://%s%s", t.Scheme, t.Host, t.Path)

	t.setConnectionOptions()

	t.Conn, _, err = t.WebsocketDialer.Dial(addr, t.RequestHeader)
	if err != nil {
		logrus.Error("dial err:", err)
		return err
	}

	return nil
}

func (t *Client) Request(reqMsgType int, reqMsg []byte) (msgType int, msg []byte, err error) {
	err = t.Conn.WriteMessage(reqMsgType, reqMsg)
	if err != nil {
		logrus.Error("WriteMessage err:", err)
		return
	}

	msgType, msg, err = t.Conn.ReadMessage()
	if err != nil {
		logrus.Error("ReadMessage err:", err)
		return
	}
	return
}

func (t *Client) RequestWithTimeout(reqMsgType int, reqMsg []byte, timeout time.Duration) (msgType int, msg []byte, err error) {
	err = t.Conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		logrus.Error("SetReadDeadline err:", err)
		return
	}

	err = t.Conn.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		logrus.Error("SetWriteDeadline err:", err)
		return
	}

	err = t.Conn.WriteMessage(reqMsgType, reqMsg)
	if err != nil {
		logrus.Error("WriteMessage err:", err)
		return
	}

	msgType, msg, err = t.Conn.ReadMessage()
	if err != nil {
		logrus.Error("ReadMessage err:", err)
		return
	}
	return
}

func (t *Client) RecvByte() (messageType int, message []byte, err error) {
	messageType, message, err = t.Conn.ReadMessage()
	if err != nil {
		logrus.Error("ReadMessage err:", err)
		return
	}
	return
}

func (t *Client) SendByte(messageType int, message []byte) (err error) {
	err = t.Conn.WriteMessage(messageType, []byte(message))
	if err != nil {
		logrus.Error("WriteMessage err:", err)
		return
	}
	return
}

func (t *Client) RecvString() (messageType int, message string, err error) {
	messageType, p, err := t.Conn.ReadMessage()
	if err != nil {
		logrus.Error("ReadMessage err:", err)
		return
	}
	message = string(p)
	return
}

func (t *Client) SendString(messageType int, message string) (err error) {
	err = t.Conn.WriteMessage(messageType, []byte(message))
	if err != nil {
		logrus.Error("WriteMessage err:", err)
		return
	}
	return
}

func (t *Client) SendText(message string) {
	err := t.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		logrus.Error("write:", err)
		return
	}
}

func (t *Client) SendBinary(data []byte) {
	err := t.Conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		logrus.Error("write:", err)
		return
	}
}

func (t *Client) DisConnect() {
	err := t.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logrus.Println("write close:", err)
		return
	}
	_ = t.Conn.Close()
}
