package network

import (
	"fmt"
	"net/http"
	"netgo/pkg/api"

	"github.com/sirupsen/logrus"
)

func NewWebsocketClient(w http.ResponseWriter, r *http.Request) (api.WebsocketClient, error) {
	return NewNhooyrWebsocketClient(w, r)
}

func NewHttpWebsocketServer(port int, endpoint string) {
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		wsClient, err := NewWebsocketClient(w, r)
		if err != nil {
			return
		}

		wsClient.ReadLoop()
	})

	logrus.Println(fmt.Sprintf("WebSocket server listening on :%d", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		logrus.Fatal("Error starting WebSocket server:", err)
	}
}
