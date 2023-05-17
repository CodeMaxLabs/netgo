package network

import (
	"context"
	"net/http"
	"netgo/pkg/api"

	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
)

type NhooyrWebsocketClient struct {
	ctx  context.Context
	conn *websocket.Conn
	w    http.ResponseWriter
	r    *http.Request
}

func NewNhooyrWebsocketClient(w http.ResponseWriter, r *http.Request) (api.WebsocketClient, error) {

	client := &NhooyrWebsocketClient{
		w:   w,
		r:   r,
		ctx: r.Context(),
	}

	err := client.Start()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (ws *NhooyrWebsocketClient) Start() error {

	conn, err := websocket.Accept(ws.w, ws.r, &websocket.AcceptOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		logrus.Println("Error accepting WebSocket connection:", err)
		return err
	}

	logrus.Println("WebSocket connection established")

	ws.conn = conn
	return nil
}

func (ws *NhooyrWebsocketClient) Stop() error {
	err := ws.conn.Close(websocket.StatusInternalError, "WebSocket connection closed")

	return err
}

func (ws *NhooyrWebsocketClient) ReadLoop() error {
	var err error
	for {
		msgType, data, err := ws.conn.Read(ws.ctx)
		if err != nil {
			logrus.Println("Error reading message:", err)

			break
		}

		if msgType == websocket.MessageText {
			logrus.Println("Received message:", string(data))
		} else if msgType == websocket.MessageBinary {
			logrus.Println("Received binary data:", data)
		}

		err = ws.conn.Write(ws.ctx, msgType, data)
		if err != nil {
			logrus.Println("Error writing message:", err)
			break
		}
	}

	return err
}
