package transport

import "qudecim/db/internal/app"

type Hub struct {
	clients map[*Client]bool

	register chan *Client

	unregister chan *Client

	app *app.App
}

func newHub(app *app.App) *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		app:        app,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		}
	}
}
