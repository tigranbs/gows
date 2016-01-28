package tree_socket

import (
	"golang.org/x/net/websocket"
	"net/http"
)

var (
	SocketHandler = websocket.Handler(socket_handler)
)

type Socket struct {
	events map[string][]func(interface{})
	ws     *websocket.Conn
}

func Listen(address, socket_path string) error {
	http.Handle(socket_path, SocketHandler)
	return http.ListenAndServe(address, nil)
}

func socket_handler(ws *websocket.Conn) {
	var (
		sock          = new(Socket)
		event_message = TreeSockEvent{}
		err           error
	)

	sock.events = make(map[string][]func(interface{}))
	sock.ws = ws
	Trigger("connection", sock)
	defer sock.Trigger("disconnect", sock)

	// Keeping connecting and receiving events
	for {
		err = websocket.JSON.Receive(sock.ws, &event_message)
		if err != nil {
			return
		}
		sock.Trigger(event_message.Event, event_message.Data)
	}
}

func (sock *Socket) Emit(event_name string, data interface{}) error {
	var (
		event_message = TreeSockEvent{}
	)
	event_message.Event = event_name
	event_message.Data = data

	return websocket.JSON.Send(sock.ws, event_message)
}
