package api

type WebsocketClient interface {
	Stop() error

	ReadLoop() error
}
